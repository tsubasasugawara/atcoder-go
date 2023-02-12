call plug#begin('~/.vim/plugged')

Plug 'scrooloose/nerdtree'
Plug 'jistr/vim-nerdtree-tabs'
Plug 't9md/vim-textmanip'
Plug 'xolox/vim-misc'
Plug 'xolox/vim-session'
Plug 'gosukiwi/vim-atom-dark'


call plug#end()

set number
set incsearch
set hidden
set backspace=indent,eol,start
set guioptions+=a
set clipboard=unnamed,autoselect
set tabstop=4
set shiftwidth=4
set title
set hlsearch
set wildmenu
set cursorline
set mouse=a
set ambiwidth=double
set history=1000
set clipboard=unnamedplus

let NERDTreeShowHidden = 1

nmap <Esc><Esc> :nohlsearch<CR><Esc>

syntax enable
colorscheme atom-dark-256 

" 行の複製
xmap <Space>d <Plug>(textmanip-duplicate-down)
nmap <Space>d <Plug>(textmanip-duplicate-down)
xmap <Space>D <Plug>(textmanip-duplicate-down)
nmap <Space>D <Plug>(textmanip-duplicate-down)

xmap <C-j> <Plug>(textmanip-move-down)
xmap <C-k> <Plug>(textmanip-move-up)
xmap <C-h> <Plug>(textmanip-move-left)
xmap <C-l> <Plug>(textmanip-move-right)

" 現在のディレクトリ直下の .vimsessions/ を取得
let s:local_session_directory = xolox#misc#path#merge(getcwd(), '.vimsession')
" 存在すれば
if isdirectory(s:local_session_directory)
  " session保存ディレクトリをそのディレクトリの設定
  let g:session_directory = s:local_session_directory
  " vimを辞める時に自動保存
  let g:session_autosave = 'yes'
  " 引数なしでvimを起動した時にsession保存ディレクトリのdefault.vimを開く
  let g:session_autoload = 'yes'
  " 1分間に1回自動保存
  let g:session_autosave_periodic = 1
else
  let g:session_autosave = 'no'
  let g:session_autoload = 'no'
endif
unlet s:local_session_directory

