package main

import (
	"strings"
)

type node struct {
	pattern  string  // 待匹配路由
	part     string  // 匹配到的一部分
	children []*node //子节点
	isWild   bool    //是否精准匹配
}

// matchChild 获取第一个匹配到的节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChild 获取匹配到的所有节点
func (n *node) matchChildren(part string) []*node {
	children := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			children = append(children, child)
		}
	}
	return children
}

// insert 插入节点
func (n *node) insert(pattern string, parts []string, height int) {
	// 1.如果是最底层节点，赋值返回
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	// 2.查找当前节点，不存在新增
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == '*' || part[0] == ':',
		}
		n.children = append(n.children, child)
	}

	// 3.继续往下写入
	child.insert(pattern, parts, height+1)
}

// search 查找节点
func (n *node) search(parts []string, height int) *node {
	// 如果是最后一层
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}

func main() {
	node := &node{
		part:     "GET",
		children: nil,
	}

	url := "hello/:name/test"
	parts := strings.Split(url, "/")
	node.insert(url, parts, 0)

	url = "hello/:name/test2"
	parts = strings.Split(url, "/")
	node.insert(url, parts, 0)

	n := node.search(parts, 0)
	if n == nil {
		panic("not found")
	}

}
