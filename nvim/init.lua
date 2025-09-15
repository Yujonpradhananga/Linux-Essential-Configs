-- bootstrap lazy.nvim, LazyVim and your plugins
require("config.lazy")
local function set_transparency()
  vim.cmd([[
hi Normal guibg=NONE ctermbg=NONE
hi NormalNC guibg=NONE ctermbg=NONE
hi SignColumn guibg=NONE ctermbg=NONE
hi StatusLine guibg=NONE ctermbg=NONE
hi StatusLineNC guibg=NONE ctermbg=NONE
hi VertSplit guibg=NONE ctermbg=NONE
hi TabLine guibg=NONE ctermbg=NONE
hi TabLineFill guibg=NONE ctermbg=NONE
hi TabLineSel guibg=NONE ctermbg=NONE
hi Pmenu guibg=NONE ctermbg=NONE
hi PmenuSel guibg=NONE ctermbg=NONE
hi NeoTreeNormal guibg=NONE ctermbg=NONE
hi NeoTreeNormalNC guibg=NONE ctermbg=NONE
hi NeoTreeWinSeparator guibg=NONE ctermbg=NONE
hi NeoTreeEndOfBuffer guibg=NONE ctermbg=NONE
hi EndOfBuffer guibg=NONE ctermbg=NONE
hi CursorLine guibg=NONE ctermbg=NONE
]])
  -- Disable cursor line highlighting
  vim.opt.cursorline = false
end

-- Set colorscheme
vim.cmd.colorscheme("wildcharm")

-- Apply transparency settings after colorscheme
set_transparency()

-- Reapply transparency on buffer enter and colorscheme change
vim.api.nvim_create_autocmd({ "BufEnter", "ColorScheme" }, {
  pattern = "*",
  callback = set_transparency,
})
