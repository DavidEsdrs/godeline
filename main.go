package main

import (
	"fmt"
	"time"

	editnode "github.com/DavidEsdrs/goditor/editNode"
	"github.com/DavidEsdrs/goditor/logger"
	text_processor "github.com/DavidEsdrs/goditor/text-processor"
)

var input string = `
A programação é uma habilidade poderosa que abre portas para infinitas possibilidades no mundo da tecnologia. Como um programador, você tem o poder de criar, moldar e otimizar /*sistemas*/, aplicativos e software de todos os tipos. É uma disciplina que combina criatividade e !~lógica~!, tornando-a <$fascinante$> para aqueles que se aventuram nesse campo.

No mundo da programação, as [[linguagens]] são as <_ferramentas_> que você utiliza para expressar suas ideias. Você mencionou que é um programador com experiência em Node.js, %-JavaScript-% e %-TypeScript-%, bem como recentemente em Golang. Essas linguagens oferecem um conjunto diversificado de *.recursos e funcionalidades.* que podem ser combinados para criar soluções incríveis.

Dominar essas linguagens é apenas o começo da jornada. À medida que você avança em sua carreira, seu interesse em matemática se tornará um trunfo valioso. A matemática está profundamente entrelaçada com muitos aspectos da programação, especialmente quando se trata de tópicos como cryptography e graphic computation.

A manipulação de **!~matrizes~! e conceitos de linear algebra se tornarão ferramentas essenciais** à medida que você mergulha mais fundo na graphic computation e na criação de jogos 2D. Essa fusão de matemática e programação resulta em criações visuais impressionantes que podem cativar o público.

À medida que você busca seu <!primeiro emprego!> na área de programação, recomendo focar em suas habilidades de fullstack para obter uma compreensão abrangente do desenvolvimento de software. Com o tempo, você pode explorar áreas como DBA (Database Administrator) e data science, que ampliarão ainda mais suas habilidades e horizontes.

Cryptography e graphics computation são campos emocionantes que exigem conhecimento especializado. Continue aprendendo e aprimorando suas habilidades, e você estará preparado para enfrentar os <!desafios!> e as oportunidades que essas áreas oferecem.

Lembre-se de que o <_Windows 10 é uma plataforma !*sólida*!_> para desenvolvimento, mas esteja aberto a outras tecnologias e sistemas operacionais à medida que sua carreira avança. O mundo da programação está {{sempre evoluindo}}, e a capacidade de se adaptar é uma habilidade fundamental.

Continue perseguindo seus objetivos com entusiasmo e determinação, pois o campo da programação está repleto de oportunidades [[emocionantes]] esperando por você!
A programação é uma habilidade poderosa que abre portas para infinitas possibilidades no mundo da tecnologia. Como um programador, você tem o poder de criar, moldar e otimizar /*sistemas*/, aplicativos e software de todos os tipos. É uma disciplina que combina criatividade e !~lógica~!, tornando-a fascinante para aqueles que se aventuram nesse campo.
`

func main() {
	l := logger.NewLogger(true)
	tree := editnode.NewEditTree()

	// for now, delimiters must be asymmetric
	tree.NewEditionType("//*", "*//")
	tree.NewEditionType("//", "//")
	tree.NewEditionType("/*", "*/")
	tree.NewEditionType("<_", "_>")
	tree.NewEditionType("!~", "~!")
	tree.NewEditionType("*.", ".*")
	tree.NewEditionType("!*", "*!")
	tree.NewEditionType("{{", "}}")
	tree.NewEditionType("<!", "!>")
	tree.NewEditionType("[[", "]]")
	tree.NewEditionType("%-", "-%")
	tree.NewEditionType("<$", "$>")

	proc := text_processor.NewProcessor(&tree, 1<<10, &l)

	start := time.Now()

	// result := proc.TokenizeText(input, true)
	result, found := proc.FoundTag("[[bom]]", 0)

	duration := time.Since(start)

	// for _, t := range result.Tokens() {
	// 	fmt.Printf("%v\n\n", t.Word)
	// }

	fmt.Printf("duration: %v\n", duration.String())
	// fmt.Printf("stats: %v tokens found\n", result.TokenQuantity)
	fmt.Printf("found token: %v\n", found)
	fmt.Printf("token: %#v token found\n", result)
}
