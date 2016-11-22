package kernel

import "github.com/eatbytes/razboy/core"

type KernelFunction func(*KernelCmd, *core.REQUEST) (*KernelCmd, error)

type KernelItem struct {
	Name string
	Fn   KernelFunction
}

type Kernel struct {
	def     KernelItem
	items   []KernelItem
	commons []string
}

var kInstance *Kernel

func Boot(def ...KernelItem) *Kernel {
	if kInstance != nil {
		return kInstance
	}

	kInstance = &Kernel{
		def: def[0],
	}

	return kInstance
}

func (k Kernel) GetDefaultItem() KernelItem {
	return k.def
}

func (k Kernel) GetItems() []KernelItem {
	return k.items
}

func (k Kernel) GetItemsName() []string {
	var names []string

	for _, item := range k.items {
		names = append(names, item.Name)
	}

	return names
}

func (k Kernel) GetCommons() []string {
	return k.commons
}

func (k *Kernel) SetDefault(item KernelItem) {
	k.def = item
}

func (k *Kernel) AddItem(item KernelItem) {
	k.items = append(k.items, item)
}
