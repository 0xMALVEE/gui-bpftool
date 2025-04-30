package main

import (
	"fmt"
	"os/user"
	"strconv"
	bpfprog "tui-bpftool/pkg/bpf-prog"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func getUsernameFromUID(uid uint32) (string, error) {
	usr, err := user.LookupId(strconv.FormatUint(uint64(uid), 10))
	if err != nil {
		return "", err
	}
	return usr.Username, nil
}

func main() {
	app := tview.NewApplication()
	app.EnableMouse(true)

	searchBpf := tview.NewInputField().
		SetLabel("Search: ").
		SetPlaceholder("Enter program name").
		SetFieldWidth(30)

	bpfListTable := tview.NewTable().SetSelectable(true, false)

	fillBpfListTable := func(programName string) {
		bpfListTable.Clear()
		bpfListTable.SetCell(0, 0, tview.NewTableCell("Name").SetExpansion(1).SetTextColor(tview.Styles.PrimaryTextColor))
		bpfListTable.SetCell(0, 1, tview.NewTableCell("Type").SetExpansion(1).SetTextColor(tview.Styles.PrimaryTextColor))
		bpfListTable.SetCell(0, 2, tview.NewTableCell("ID").SetExpansion(1).SetTextColor(tview.Styles.PrimaryTextColor))
		bpfListTable.SetCell(0, 3, tview.NewTableCell("Created By").SetExpansion(1).SetTextColor(tview.Styles.PrimaryTextColor))

		ebpgProgList, _ := bpfprog.GetProgListWithInfo(programName)

		for i, prog := range ebpgProgList {
			uid, _ := prog.ProgramInfo.CreatedByUID()
			username, _ := getUsernameFromUID(uid)
			bpfListTable.SetCell(i+1, 0, tview.NewTableCell(prog.ProgramInfo.Name).SetExpansion(1).SetTextColor(tview.Styles.PrimaryTextColor))
			bpfListTable.SetCell(i+1, 1, tview.NewTableCell(prog.Type).SetExpansion(1).SetTextColor(tview.Styles.PrimaryTextColor))
			bpfListTable.SetCell(i+1, 2, tview.NewTableCell(fmt.Sprintf("%d", prog.ProgramInfo.ID)).SetExpansion(1).SetTextColor(tview.Styles.PrimaryTextColor))
			bpfListTable.SetCell(i+1, 3, tview.NewTableCell(username).SetExpansion(1).SetTextColor(tview.Styles.PrimaryTextColor))
		}
	}

	fillBpfListTable("")

	bpfListTable.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			row, col := bpfListTable.GetSelection()
			n := bpfListTable.GetCell(row, col).Text
			fmt.Println("clicked enter", row, col, n)
		}
		return event
	})

	searchBpf.SetChangedFunc(func(text string) {
		if text == "" {
			fillBpfListTable("")
			return
		}
		fillBpfListTable(text)
	})

	layout := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(searchBpf, 1, 0, false).AddItem(bpfListTable, 0, 1, true)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			if bpfListTable.HasFocus() {
				app.SetFocus(searchBpf)
				return nil
			}
			app.SetFocus(bpfListTable)
			return nil
		}
		return event
	})

	if err := app.SetRoot(layout, true).SetFocus(searchBpf).Run(); err != nil {
		panic(err)
	}
}
