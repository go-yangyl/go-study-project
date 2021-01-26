package gin

import (
	"strings"
)

// 路由节点  /p/c/:lang
type node struct {
	pattern  string  // 待匹配路由，全路径
	part     string  // 路由中的一部分
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWaild  bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

// 不同的路由加入node节点
func (n *node) Insert(pattern string, parts []string, height int) {
	// 匹配完毕，递归中止
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.MatchInsert(part)

	if child == nil {
		child = &node{part: part, isWaild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.Insert(pattern, parts, height+1)
}

func (n *node) MatchInsert(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWaild {
			return child
		}
	}
	return nil
}
func (n *node) Search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.MatchSearch(part)

	for _, child := range children {
		result := child.Search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

func (n *node) MatchSearch(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWaild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 解析请求url
func ParsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}
