package main

import (
	"fmt"
	"time"

	editnode "github.com/DavidEsdrs/goditor/editNode"
	"github.com/DavidEsdrs/goditor/logger"
	text_processor "github.com/DavidEsdrs/goditor/text-processor"
)

// holds the maximum window length in which we go search for a closing tag
const maxBufferLength = 1 << 12

var input string = `
<html>
	<head>
		<title>Página Sem Atributos</title>
	</head>
	<body>
		<header>
			<h1>Exemplo de Página HTML</h1>
			<nav>
				<ul>
					<li><a>Início</a></li>
					<li><a>Sobre</a></li>
					<li><a>Contato</a></li>
				</ul>
			</nav>
		</header>
		<main>
			<section>
				<h2>Seção 1</h2>
				<p>Esta é a primeira seção do conteúdo.</p>
			</section>
			<section>
				<h2>Seção 2</h2>
				<p>Esta é a segunda seção do conteúdo.</p>
				<ul>
					<li>Item 1</li>
					<li>Item 2</li>
					<li>Item 3</li>
				</ul>
			</section>
		</main>
		<footer>
			<p>Rodapé da página.</p>
		</footer>
	</body>
</html>
`

func main() {
	l := logger.NewLogger(false)
	tree := editnode.NewEditTree()

	// for now, delimiters must be asymmetric
	tree.NewEditionType("<html>", "</html>")
	tree.NewEditionType("<head>", "</head>")
	tree.NewEditionType("<body>", "</body>")
	tree.NewEditionType("<header>", "</header>")
	tree.NewEditionType("<h1>", "</h1>")
	tree.NewEditionType("<nav>", "</nav>")
	tree.NewEditionType("<ul>", "</ul>")
	tree.NewEditionType("<li>", "</li>")
	tree.NewEditionType("<a>", "</a>")
	tree.NewEditionType("<main>", "</main>")
	tree.NewEditionType("<section>", "</section>")
	tree.NewEditionType("<h2>", "</h2>")
	tree.NewEditionType("<p>", "</p>")
	tree.NewEditionType("<footer>", "</footer>")

	proc := text_processor.NewProcessor(&tree, maxBufferLength, &l)

	l.LogProcessf("Starting tokenization...")
	start := time.Now()

	result := proc.Tokenize(input, false)

	duration := time.Since(start)
	l.LogProcessf("Ending tokenization")

	for _, t := range result.Tokens() {
		l.LogProcessf("%+v\n\n", t)
	}

	fmt.Printf("duration: %v\n", duration.String())
	fmt.Printf("stats: %v tokens found\n", result.TokenQuantity)
}
