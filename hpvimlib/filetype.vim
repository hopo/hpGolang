"au BufNewFile,BufRead *html,*.gohtml source ~/.vim/syntax/gohtml.vim
"autocmd BufNewFile,BufRead *.gohtml setfiletype gohtml


" *** my filetype file
if exists("did_load_filetypes")
    finish
endif

augroup filetypedetect
    autocmd! BufEnter *.cpp setfiletype cpp
	autocmd! BufWinLeave *.cpp	setfiletype u_cpp
augroup END

augroup filetypedetect
    autocmd! BufEnter *.go setfiletype go
	autocmd! BufWinLeave *.go	setfiletype u_go
augroup END

augroup filetypedetect
    autocmd! BufEnter *.gohtml	setfiletype gohtml
	autocmd! BufWinLeave *.gohtml	setfiletype u_gohtml
augroup END

augroup filetypedetect
    autocmd! BufEnter *.py	setfiletype python
	"autocmd! BufWinLeave *.py	setfiletype u_gohtml
augroup END
