package kernel

import (
	"errors"

	"github.com/eatbytes/razboy"
)

func (k Kernel) GetDefaultItem() *KernelItem {
	return k.def
}

func (k Kernel) GetItems() []*KernelItem {
	return k.items
}

func (k Kernel) GetFormerCmd() *KernelCmd {
	return k.former
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

func (k *Kernel) SetFormerCmd(kc *KernelCmd) {
	k.former = kc
}

func (k *Kernel) SetDefault(item *KernelItem) {
	k.def = item
}

func (k *Kernel) SetItems(items []*KernelItem) {
	k.items = items
}

func (k *Kernel) AddItem(item *KernelItem) {
	k.items = append(k.items, item)
}

func (k *Kernel) Stop() {
	k.run = false
}

func (k *Kernel) UpdatePrompt(url, scope string) {
	if k.readline == nil {
		return
	}

	k.readline.SetPrompt("(" + url + "):" + scope + "$ ")
}

func KernelDefault(kc *KernelCmd, config *razboy.Config) (*KernelCmd, error) {
	return kc, errors.New("No default fonction defined")
}
