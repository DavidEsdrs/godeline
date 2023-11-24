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

	proc := godeline.NewProcessor(&tree, 1<<12)

	_, err := proc.Tokenize(input)

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

	proc := godeline.NewProcessor(&tree, 1<<12)

	proc.Sanitize()

	result, err := proc.Tokenize(input)

	result.Tokens() // array de tokens

	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestNested(t *testing.T) {
	tree := editnode.NewEditTree()
	tree.AddDelimiterType("[[", "]]")
	proc := godeline.NewProcessor(&tree, 1<<12)
	_, err := proc.Tokenize("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec iaculis mauris in fringilla ornare. In hac habitasse platea dictumst. Mauris et nisi id turpis vulputate fermentum. Nulla ac aliquet lorem, sed placerat est. Nam lacinia est ac velit malesuada, non lacinia erat laoreet. Praesent purus diam, consequat quis libero at, [[vu[[putate]] idawdawd ]]mperdiet purus. In ut luctus ex. Aenean semper non orci blandit varius. Vestibulum lectus est, cursus a orci sed, gravida pretium massa. Nam id elit quis massa aliquet tristique. Ut porttitor aliquam semper. Nam arcu ipsum, aliquam quis aliquam a, varius sed nulla. Cras at ante eu libero dapibus iaculis et id urna. Donec ac velit tellus. Duis vestibulum nec mi et vehicula. Pellentesque ornare volutpat rhoncus.")
	if err != nil {
		t.Error(err)
	}
}

func TestErrorClosingNotFound(t *testing.T) {
	tree := editnode.NewEditTree()
	tree.AddDelimiterType("[[", "]]")
	proc := godeline.NewProcessor(&tree, 1<<12)
	proc.StopOnError()
	_, err := proc.Tokenize("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec iaculis mauris in fringilla ornare. In hac habitasse platea dictumst. Mauris et nisi id turpis vulputate fermentum. Nulla ac aliquet lorem, sed placerat est. Nam lacinia est ac velit malesuada, non lacinia erat laoreet. Praesent purus diam, consequat quis libero at, [[vuputate]] idawdawd [[ mperdiet purus. In ut luctus ex. Aenean semper non orci blandit varius. Vestibulum lectus est, cursus a orci sed, gravida pretium massa. Nam id elit quis massa aliquet tristique. Ut porttitor aliquam semper. Nam arcu ipsum, aliquam quis aliquam a, varius sed nulla. Cras at ante eu libero dapibus iaculis et id urna. Donec ac velit tellus. Duis vestibulum nec mi et vehicula. Pellentesque ornare volutpat rhoncus.")

	if err == nil || err != godeline.ErrClosingTagNotFound {
		t.Errorf("failed - error didn't happen as expected")
	}
}
