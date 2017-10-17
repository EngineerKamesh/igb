package carsdemo

import (
	"bytes"
	"encoding/gob"
	"html/template"

	"github.com/EngineerKamesh/igb/igweb/shared/models"

	"honnef.co/go/js/dom"
	"honnef.co/go/js/xhr"
)

const CarItemTemplate = `
	<td class="mdl-data-table__cell--non-numeric">{{.ModelName}}</td>
	<td class="mdl-data-table__cell--non-numeric">{{.Color}}</td>
	<td class="mdl-data-table__cell--non-numeric">{{.Manufacturer}}</td>
`

var D = dom.GetWindow().Document()

func InitializePage() {

	cars()

}

func cars() {

	nano := models.Car{ModelName: "Nano", Color: "Yellow", Manufacturer: "Tata"}
	ambassador := models.Car{ModelName: "Ambassador", Color: "White", Manufacturer: "HM"}
	omni := models.Car{ModelName: "Omni", Color: "Red", Manufacturer: "Maruti Suzuki"}
	cars := []models.Car{nano, ambassador, omni}

	autoTableBody := D.GetElementByID("autoTableBody")
	for i := 0; i < len(cars); i++ {
		trElement := D.CreateElement("tr")
		tpl := template.New("template")
		tpl.Parse(CarItemTemplate)
		var buff bytes.Buffer
		tpl.Execute(&buff, cars[i])
		trElement.SetInnerHTML(buff.String())
		autoTableBody.AppendChild(trElement)
	}

	var carsDataBuffer bytes.Buffer
	enc := gob.NewEncoder(&carsDataBuffer)
	enc.Encode(cars)

	xhrResponse, err := xhr.Send("POST", "/cars-data", carsDataBuffer.Bytes())

	if err != nil {
		println(err)
	}

	println("xhrResponse: ", string(xhrResponse))

}
