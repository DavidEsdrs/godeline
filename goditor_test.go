package goditor_test

import (
	"testing"

	"github.com/DavidEsdrs/goditor"
	editnode "github.com/DavidEsdrs/goditor/editNode"
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
	tree.NewEditionType("<footer>", "</footer>")

	proc := goditor.NewProcessor(&tree, 1<<12, nil)

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
	tree.NewEditionType("<footer>", "</footer>")

	proc := goditor.NewProcessor(&tree, 1<<12, nil)

	_, err := proc.Tokenize(input, true)

	if err != nil {
		t.Errorf(err.Error())
	}
}
