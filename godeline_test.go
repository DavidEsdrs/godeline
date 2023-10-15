package godeline_test

import (
	"testing"

	"github.com/DavidEsdrs/godeline"
	editnode "github.com/DavidEsdrs/godeline/edit-node"
)

func TestTokenize(t *testing.T) {
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

	tree := editnode.NewEditTree()

	tree.AddDelimiterType("<html>", "</html>")
	tree.AddDelimiterType("<head>", "</head>")
	tree.AddDelimiterType("<body>", "</body>")
	tree.AddDelimiterType("<header>", "</header>")
	tree.AddDelimiterType("<h1>", "</h1>")
	tree.AddDelimiterType("<nav>", "</nav>")
	tree.AddDelimiterType("<ul>", "</ul>")
	tree.AddDelimiterType("<li>", "</li>")
	tree.AddDelimiterType("<a>", "</a>")
	tree.AddDelimiterType("<main>", "</main>")
	tree.AddDelimiterType("<section>", "</section>")
	tree.AddDelimiterType("<h2>", "</h2>")
	tree.AddDelimiterType("<p>", "</p>")
	tree.AddDelimiterType("<footer>", "</footer>")
	tree.AddDelimiterType("<footer>", "</footer>")

	proc := godeline.NewProcessor(&tree, 1<<12, nil)

	_, err := proc.Tokenize(input, false)

	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestTokenizeAndSanitization(t *testing.T) {
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

	tree := editnode.NewEditTree()

	tree.AddDelimiterType("<html>", "</html>")
	tree.AddDelimiterType("<head>", "</head>")
	tree.AddDelimiterType("<body>", "</body>")
	tree.AddDelimiterType("<header>", "</header>")
	tree.AddDelimiterType("<h1>", "</h1>")
	tree.AddDelimiterType("<nav>", "</nav>")
	tree.AddDelimiterType("<ul>", "</ul>")
	tree.AddDelimiterType("<li>", "</li>")
	tree.AddDelimiterType("<a>", "</a>")
	tree.AddDelimiterType("<main>", "</main>")
	tree.AddDelimiterType("<section>", "</section>")
	tree.AddDelimiterType("<h2>", "</h2>")
	tree.AddDelimiterType("<p>", "</p>")
	tree.AddDelimiterType("<footer>", "</footer>")
	tree.AddDelimiterType("<footer>", "</footer>")

	proc := godeline.NewProcessor(&tree, 1<<12, nil)

	_, err := proc.Tokenize(input, true)

	if err != nil {
		t.Errorf(err.Error())
	}
}
