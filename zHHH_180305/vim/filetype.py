
"
" *** User filetype file


if exists("did_load_filetypes")
	finish
endif

augroup filetypedetect
	autocmd! BufEnter *.h,*.c,*.cpp setfiletype c
	autocmd! BufWinLeave *.h,*.c,*.cpp setfiletype u_c
augroup END

augroup filetypedetect
	autocmd! BufEnter *.py setfiletype python
	autocmd! BufWinLeave *.py setfiletype u_python
augroup END

augroup filetypedetect
	autocmd! BufEnter *.go setfiletype go
	autocmd! BufWinLeave *.go setfiletype u_go
augroup END
