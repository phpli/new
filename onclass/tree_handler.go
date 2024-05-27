package main

import "strings"

type HandlerBasedOnTree struct {
	root *node
}

type node struct {
	path     string
	children []*node
	//如果这是叶子节点
	//那么匹配上之后页可以调用该方法
	handler handlerFunc
}

func (h *HandlerBasedOnTree) ServeHTTP(c *Context) {
	//TODO implement me
	panic("implement me")
}

func (h *HandlerBasedOnTree) Route(method string, pattern string, handlerFunc func(ctx *Context)) {
	//TODO implement me
	pattern = strings.Trim(pattern, "/")
	//panic("implement me")
	//[user,friends]
	paths := strings.Split(pattern, "/")
	cur := h.root
	for index, path := range paths {
		matchChild, ok := h.findMatchChild(cur, path)
		if ok {
			cur = matchChild
		} else {
			h.createSubTree(cur, paths[index:], handlerFunc)
			return
		}
	}
}

func (h *HandlerBasedOnTree) findMatchChild(root *node,
	path string) (*node, bool) {
	var wildcarNode *node
	for _, child := range root.children {
		if child.path == path && child.path != "*ß" {
			return child, true
		}
		if child.path == "*" {
			wildcarNode = child
		}
	}
	return wildcarNode, wildcarNode != nil
}

func (h *HandlerBasedOnTree) createSubTree(root *node,
	paths []string, handlerFunc handlerFunc) {
	cur := root
	for _, path := range paths {
		nn := newNode(path)
		cur.children = append(cur.children, nn)
		cur = nn
	}
}

func newNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 0, 8),
	}
}

func (h *HandlerBasedOnTree) findRouter(path string) (handlerFunc, bool) {
	paths := strings.Split(strings.Trim(path, "/"), "/")
	cur := h.root
	for _, path := range paths {
		matchChild, ok := h.findMatchChild(cur, path)
		if !ok {
			return nil, false
		}
		cur = matchChild
	}
	if cur.handler == nil {
		return nil, false
	}
	return cur.handler, true
}
