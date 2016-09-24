package utils

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/satori/go.uuid"
)

func pdfGenerator() {
	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)
	fmt.Println(u1)

	u1 = uuid.NewV5(
		uuid.UUID{},
		"lala",
	)
	fmt.Printf("UUIDv5: %s\n", u1)
	fmt.Println(u1)

	// Parsing UUID from string input
	id := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	u2, err := uuid.FromString(id)
	if err != nil {
		fmt.Printf("Something gone wrong: %s", err)
	}
	fmt.Printf("Successfully parsed: %s", u2)
}

// Natural ...
func Natural(brandName string, userID int) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	write := func(str string) {
		tr := pdf.UnicodeTranslatorFromDescriptor("")
		pdf.MultiCell(170, 5, tr(str), "", "", false)
		pdf.Ln(5)
	}
	pdf.SetHeaderFunc(func() {
		pdf.SetFont("Arial", "B", 15)
		pdf.Cell(80, 5, "")
		pdf.CellFormat(30, 10, "PODER ESPECIAL", "", 2, "C", false, 0, "")
		pdf.Ln(20)
	})
	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.SetMargins(4, 5, 0)
	pdf.SetFont("Arial", "", 10)

	pdf.Line(10, 30, 200, 30)
	pdf.Ln(0)
	pdf.Cell(15, 0, "")

	var text = ""
	name := "Rodrigo"
	lastname := "Fuenzalida"
	rut := "16.642.817-3"
	calle := "Avenida Fleming 9455"
	comuna := "Las Condes"
	ciudad := "Santiago"
	pais := "Chile"
	// // nombre, apellido, rut, calle, comuna, ciudad, pais
	text = fmt.Sprintf("Don %s %s, C.N.I. N° %s, domiciliado en %s, %s y ciudad de %s de %s, otorga por el presente instrumento poder especial a los señores Roberto Fasani Puelma y José Rivera Rojas, ambos abogados habilitados para el ejercicio de la profesión, domiciliados en calle Nueva York Nº 9, Piso 14, Santiago, de esta Ciudad, poder tan amplio y suficiente como sea necesario conforme a Derecho, para que representen en Chile y/o en otros países, ante todos los tribunales, entidades y autoridades administrativas o judiciales que correspondan, en cualquier solicitud, gestión o litigio relacionado directa o indirectamente con derechos de propiedad intelectual e industrial, incluyendo patentes, modelos de utilidad, diseños industriales, dibujos industriales, esquemas de trazado, marcas comerciales, nombres comerciales, frases de propaganda, signos distintivos, indicaciones geográficas, denominaciones de origen, nombres de dominio de internet, variedades vegetales, derechos de autor y derechos conexos.", name, lastname, rut, calle, comuna, ciudad, pais)

	write(text)
	pdf.Cell(15, 0, "")

	text = "Conforme lo anterior, los referidos mandatarios, actuando individual o conjuntamente, ya sea ellos mismos o a través de terceros con delegación de poder, podrán actuar con las más amplias y suficientes facultades para cumplir las gestiones indicadas; podrán requerir de las oficinas, entidades y autoridades pertinentes, el registro o renovación de cualquiera y/o de todas sus marcas comerciales, nombres comerciales, frases de propaganda, signos distintivos, nombres de dominio de internet, patentes, modelos de utilidad, diseños industriales, variedades vegetales o derechos de autor y/o conexos estando, con tal propósito, facultados para efectuar ante las mencionadas oficinas, entidades y autoridades, todas las gestiones necesarias, a saber, presentar solicitudes y especificaciones; formular declaraciones; deducir recursos y reclamos; formular y firmar descripciones y enmiendas; modificar, agregar y/o suprimir reivindicaciones; otorgar y recibir cesiones; pagar todos los impuestos, derechos y cualquier otro pago determinado por la ley y retirar los mismos de ser necesario; recibir toda clase de documentos y valores; hacer modificaciones en todos los documentos presentados; solicitar testimonios; desistirse de procedimientos; solicitar copias autorizadas; solicitar certificaciones de cualquier tipo; prestar declaraciones juradas; ceder solicitudes y aceptar cesiones de solicitudes, transferir, aceptar transferencias y suscribir cualquier tipo de contratos relativos a derechos de propiedad intelectual y/o industrial; limitar y desistirse en todo o en parte de las solicitudes; solicitar la cancelación voluntaria de todo o parte de los registros concedidos; solicitar la aprobación y/o inscripción de contratos de licencias, franchising, transferencias de tecnología, transferencias o cesiones de derechos de propiedad intelectual y/o industrial y cualesquiera otros contratos; solicitar la inscripción de fusiones, cambios de nombres, embargos, prendas y todo tipo de medidas precautorias; oponerse y protestar contra cualquiera solicitud o registro que, a juicio del apoderado, pudieran prestarse a confusión o infringir y/o perjudicar de cualquier otro modo las marcas comerciales, nombres comerciales, frases de propaganda, signos distintivos, nombres de dominio de internet, patentes, modelos de utilidad, diseños industriales, variedades vegetales o derechos de autor y/o derechos conexos del poderdante, con facultad, asimismo, a su discreción, de renunciar o no a acciones judiciales, recursos o plazos legales, transigir judicial y extrajudicialmente, desistirse de los recursos o de la acción deducida, absolver posiciones, someter a arbitraje y conferir a los árbitros las facultades de arbitradores; percibir; designar patrocinantes; revocar patrocinios y poderes anteriores; ratificar lo obrado por ellos o por otras personas; representar al poderdante en juicio; iniciar, proseguir y contestar acciones de oposición, nulidad, caducidad, así como actuar ante los tribunales administrativos, judiciales, ordinarios y arbitrales, con facultad para entablar toda clase de acciones y recursos civiles, comerciales y criminales; presentar denuncias y/o querellas criminales y ejercer la defensa judicial por falsificación, imitación, utilización indebida y cualquier otra infracción relacionada con las materias previamente enunciadas; desistir de las acciones deducidas; contestar y aceptar demandas; renunciar a los recursos y los términos legales; celebrar salidas alternativas; presentar apelaciones, reposiciones, nulidades y cualquier otro recurso legal o reglamentario."
	//
	write(text)
	pdf.Cell(15, 0, "")
	text = "Para el evento de cualquier dificultad respecto de la ejecución, aplicación y cumplimiento del presente mandato, las partes fijan su domicilio en la ciudad de Santiago y se someten a la jurisdicción y competencia de sus Tribunales de Justicia.-"
	//
	write(text)
	pdf.Cell(15, 0, "")
	pdf.Ln(15)
	//
	pdf.Cell(95, 20, "")
	pdf.Write(1, "________________________")
	pdf.Ln(5)
	pdf.Cell(112, 20, "")
	pdf.SetFont("Arial", "B", 10)
	pdf.Write(1, " FIRMA")
	pdf.SetFont("Arial", "", 10)
	pdf.Ln(5)
	pdf.Cell(105, 20, "")
	pdf.Write(1, fmt.Sprintf("%s %s", name, lastname))
	pdf.Ln(5)
	pdf.Cell(108, 20, "")
	pdf.Write(1, rut)

	pdf.Line(10, 30, 200, 30)
	pdf.Ln(20)

	// $pdf->Output('pdfs/'.$brand_name.'_'.$nombre.'_'.$apellido.'.pdf', 'F');
	filename := fmt.Sprintf("%s_%s_%s.pdf", brandName, name, lastname)
	// path := fmt.Sprintf("/tmp/pdfs/%d/%s/%s", userID, brandName, filename)
	if err := pdf.OutputFileAndClose(filename); err != nil {
		fmt.Println("Error while creating pdf")
		fmt.Println(err.Error())
	}
}
