package template

var (
	TexTemplate = `\documentclass[11pt]{jarticle}
\usepackage[senior]{fireport}

\begin{document}
\thispagestyle{empty}

\section*{題目一覧}

\begin{table}[htb]
\begin{tabular}{|c|c|p{31em}|}

{{.Table}}

\end{tabular}
\end{table}

\end{document}
`

	B3Template = `
\multicolumn{3}{l}{☆学部3年生} \\
\hline
学籍番号 & 氏名 & 題目 \\ \hline
{{.B3}}\multicolumn{3}{l}{} \\`

	B4Template = `
\multicolumn{3}{l}{☆学部4年生} \\
\hline
学籍番号 & 氏名 & 題目 \\ \hline
{{.B4}}\multicolumn{3}{l}{} \\`

	M1Template = `
\multicolumn{3}{l}{☆大学院1年生} \\
\hline
学籍番号 & 氏名 & 題目 \\ \hline
{{.M1}}\multicolumn{3}{l}{} \\`
)
