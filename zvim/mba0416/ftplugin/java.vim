"
" *** java filetype

noremap gt :!echo ' [ *** java compile *** ] ' && javac % -d ~/hpWorks/javatest && ~/hpWorks/javatest/ && java 

inoremap main<TAB> public static void main(String[] args) 
inoremap println<TAB> System.out.println
