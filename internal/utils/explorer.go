package utils

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"
	"strings"
)

// Node 树节点
type Node struct {
	Name     string  `json:"name"`     // 目录或文件名
	ShowName string  `json:"showName"` // 目录或文件名（不包含后缀）
	Path     string  `json:"path"`     // 目录或文件完整路径
	Link     string  `json:"link"`     // 文件访问URI
	Active   string  `json:"active"`   // 当前活跃的文件
	Children []*Node `json:"children"` // 目录下的文件或子目录
	IsDir    bool    `json:"isDir"`    // 是否为目录 true: 是目录 false: 不是目录
}

// Option 遍历选项
type Option struct {
	RootPath   []string `yaml:"rootPath"`   // 目标根目录
	SubFlag    bool     `yaml:"subFlag"`    // 遍历子目录标志 true: 遍历 false: 不遍历
	IgnorePath []string `yaml:"ignorePath"` // 忽略目录
	IgnoreFile []string `yaml:"ignoreFile"` // 忽略文件
}

// 当前再循环的Dir路径
var CurDirPath string

// Explorer 遍历多个目录
//
//	option : 遍历选项
//	tree : 遍历结果
func Explorer(option Option) (Node, error) {
	// 根节点
	var root Node

	// 多个目录搜索
	for _, p := range option.RootPath {
		// 空目录跳过
		if strings.TrimSpace(p) == "" {
			continue
		}

		var child Node

		// 目录路径
		CurDirPath = p
		child.Path = p

		// 递归
		explorerRecursive(&child, &option)

		root.Children = append(root.Children, &child)
	}

	return root, nil
}

// 递归遍历目录
//
//	node : 目录节点
//	option : 遍历选项
func explorerRecursive(node *Node, option *Option) {
	// 节点的信息
	p, err := os.Stat(node.Path)
	if err != nil {
		log.Println(err)
		return
	}
	// 是否为目录
	node.IsDir = p.IsDir()

	// 非目录，返回
	if !p.IsDir() {
		return
	}

	// 目录中的文件和子目录
	sub, err := ioutil.ReadDir(node.Path)
	if err != nil {
		info := "目录不存在，或打开错误。"
		log.Printf("%v: %v", info, err)
		return
	}

	var mdFiles []*Node

	for _, f := range sub {
		tmp := path.Join(node.Path, f.Name())
		var child Node
		// 完整子目录
		child.Path = tmp
		// 目录（或文件）名
		child.Name = f.Name()
		if !f.IsDir() {
			// 访问路径
			child.Link = CustomURLEncode(strings.TrimPrefix(strings.TrimSuffix(tmp, path.Ext(f.Name())), CurDirPath))

			// 目录或文件名（不包含后缀）
			child.ShowName = strings.TrimSuffix(f.Name(), path.Ext(f.Name()))
		} else {
			// 如果是目录，无需处理后缀
			child.Link = CustomURLEncode(strings.TrimPrefix(tmp, CurDirPath))
			child.ShowName = f.Name()
		}
		if strings.Index(child.ShowName, "@") != -1 {
			child.ShowName = child.ShowName[strings.Index(child.ShowName, "@")+1:]
		}
		// 是否为目录
		child.IsDir = f.IsDir()

		// 目录
		if f.IsDir() {
			// 查找子目录
			if option.SubFlag {
				// 不在忽略目录中的目录，进行递归查找
				if !IsInSlice(option.IgnorePath, f.Name()) {
					node.Children = append(node.Children, &child)
					explorerRecursive(&child, option)
				}
			}
		} else { // 文件
			// 过滤非md文件
			if path.Ext(f.Name()) != ".md" {
				continue
			}

			// 非忽略文件，添加到结果中
			if IsInSlice(option.IgnoreFile, f.Name()) {
				continue
			}

			mdFiles = append(mdFiles, &child)
		}
	}

	// 超过指定数量的.md文件进行随机展示
	if len(mdFiles) > 150 {
		rand.Shuffle(len(mdFiles), func(i, j int) {
			mdFiles[i], mdFiles[j] = mdFiles[j], mdFiles[i]
		})
		node.Children = append(node.Children, mdFiles[:150]...)
	} else {
		node.Children = append(node.Children, mdFiles...)
	}
}
