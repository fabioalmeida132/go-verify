package utils

import (
	"context"
	"encoding/json"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/go-verify/pkg/models"
	"log"
	"time"
)

func Ofac(name string) []models.Ofac {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	jsText := jsGetText("#gvSearchResults > tbody")

	var res string
	var scrappeds []models.Ofac

	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://sanctionssearch.ofac.treas.gov`),
		chromedp.WaitVisible(`body`),
		chromedp.SendKeys(`#ctl00_MainContent_txtLastName`, name),
		chromedp.SetValue(`#ctl00_MainContent_Slider1_Boundcontrol`, "80"),
		chromedp.Click(`#ctl00_MainContent_btnSearch`),
		chromedp.Sleep(3*time.Second),
		chromedp.Nodes(`#gvSearchResults > tbody`, &nodes, chromedp.AtLeast(0)),
	)
	if err != nil {
		log.Fatal(err)
	}

	if len(nodes) > 0 {
		err = chromedp.Run(ctx, chromedp.WaitVisible(`#gvSearchResults > tbody`),
			chromedp.Evaluate(jsText, &res),
		)

		err = json.Unmarshal([]byte(res), &scrappeds)
		if err != nil {
			log.Fatal(err)
		}
	}
	chromedp.Stop()
	return scrappeds
}

func jsGetText(sel string) (js string) {
	//jsText := jsGetText("#gvSearchResults > tbody > tr > td > a")
	const funcJS = `function getText(sel) {
				var text = "";
				var elements = document.body.querySelectorAll(sel);
                var current = elements[0];
				for (var i = 0; i < current.children.length; i++) {
					text += JSON.stringify({ 
						name: current.children[i].cells[0].innerText,
						link: current.children[i].cells[0].children[0].href,
						program: current.children[i].cells[3].innerText,
					 	list: current.children[i].cells[4].innerText,
					 	score: current.children[i].cells[5].innerText
					});
					if (i < current.children.length - 1) {text += ",";}
				}
				return "[" + text + "]";
			 };`

	invokeFuncJS := `var a = getText('` + sel + `'); a;`

	return funcJS + invokeFuncJS
}
