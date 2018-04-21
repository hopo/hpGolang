"
" *** java filetype

noremap gr :!echo " >>* java compile & run *<< " && javac % -d ~/hpWorkspace/javatest && cd ~/hpWorkspace/javatest && java 

noremap gt :!echo " >>* java compile *<< " && javac % -d ~/hpWorkspace/javatest<CR>
