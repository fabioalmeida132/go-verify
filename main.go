package main

import (
	"context"
	"log"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type Scrapped struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

func main() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	name := ""
	link := ""
	rows := []*cdp.Node{}

	// scrappeds := []Scrapped{}

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://sanctionssearch.ofac.treas.gov`),
		chromedp.WaitVisible(`body`),
		chromedp.SendKeys(`#ctl00_MainContent_txtLastName`, "beiramar"),
		chromedp.SetValue(`#ctl00_MainContent_Slider1_Boundcontrol`, "80"),
		chromedp.Click(`#ctl00_MainContent_btnSearch`),
		chromedp.WaitVisible(`#gvSearchResults > tbody`),
		chromedp.Evaluate(`document.querySelector("#gvSearchResults > tbody > tr > td > a").innerText`, &name),
		chromedp.Evaluate(`document.querySelector("#gvSearchResults > tbody > tr > td > a").href`, &link),
		chromedp.Nodes(`#gvSearchResults > tbody > tr`, &rows, chromedp.ByQueryAll),

		// chromedp.ActionFunc(func(ctx context.Context) error {
		// 	for _, row := range rows {
		// 		var name string
		// 		err := chromedp.Run(ctx,
		// 			chromedp.Text(row, &name, chromedp.ByQuery),
		// 		)
		// 		if err != nil {
		// 			return err
		// 		}
		// 		log.Println(name)
		// 	}
		// 	return nil
		// }),
		chromedp.Stop(),
	)
	if err != nil {
		log.Fatal(err)
	}

}
