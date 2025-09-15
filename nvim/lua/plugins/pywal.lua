-- ~/.config/nvim/lua/plugins/pywal.lua
return {
  {
    "dylanaraps/wal.vim",
    lazy = false,
    priority = 1000,
    config = function()
      vim.cmd("colorscheme wal")
    end,
  },
}
