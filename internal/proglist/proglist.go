package proglist

import (
	"fmt"
	"os/user"
	"strconv"

	"tui-bpftool/internal"
	"tui-bpftool/internal/app"

	bpfprog "tui-bpftool/pkg/bpf-prog"

	"github.com/cilium/ebpf"
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

func GetProgListView(app *app.Application) *tview.Flex {
	tviewApp := app.App
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
			id, _ := prog.ProgramInfo.ID()
			bpfListTable.SetCell(i+1, 0, tview.NewTableCell(prog.ProgramInfo.Name).SetExpansion(1).SetTextColor(tview.Styles.PrimaryTextColor))
			bpfListTable.SetCell(i+1, 1, tview.NewTableCell(prog.Type).SetExpansion(1).SetTextColor(tview.Styles.PrimaryTextColor))
			bpfListTable.SetCell(i+1, 2, tview.NewTableCell(fmt.Sprintf("%d", id)).SetExpansion(1).SetTextColor(tview.Styles.PrimaryTextColor))
			bpfListTable.SetCell(i+1, 3, tview.NewTableCell(username).SetExpansion(1).SetTextColor(tview.Styles.PrimaryTextColor))
		}
	}

	fillBpfListTable("")

	bpfListTable.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			row, _ := bpfListTable.GetSelection()
			n := bpfListTable.GetCell(row, 2).Text
			progID, _ := strconv.Atoi(n)
			app.SelectedProgramID = ebpf.ProgramID(progID)
			app.SetCurrentView(func() {
				app.CurrentView = internal.ProgramMapsView
				app.App.SetRoot(app.MapsView, true)
			})
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

	tviewApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			if bpfListTable.HasFocus() {
				tviewApp.SetFocus(searchBpf)
				return nil
			}
			tviewApp.SetFocus(bpfListTable)
			return nil
		}
		return event
	})
	return layout
}
