
import subprocess
def read_xresources(prefix):
    props = {}
    #x = subprocess.run(['xrdb', '-query'], capture_output=True, check=True, text=True)
    lines = x.stdout.split('\n')
    for line in filter(lambda l : l.startswith(prefix), lines):
        prop, _, value = line.partition(':\t')
        props[prop] = value
    return props

config.load_autoconfig()  # or config.load_autoconfig(False) if you don't want GUI settings
import gruvbox
gruvbox.apply_theme(c)

c.tabs.show = "multiple"
c.tabs.width="25%"
c.tabs.title.format = "{audio}{current_title}"
c.tabs.padding={"bottom": 5, "left": 5, "right": 5, "top": 5}
c.fonts.web.size.default = 20
c.fonts.web.size.minimum = 12

c.url.searchengines = {
# note - if you use duckduckgo, you can make use of its built in bangs, of which there are many! https://duckduckgo.com/bangs
        'DEFAULT': 'https://duckduckgo.com/?q={}',
        '!aw': 'https://wiki.archlinux.org/?search={}',
        '!apkg': 'https://archlinux.org/packages/?sort=&q={}&maintainer=&flagged=',
        '!gh': 'https://github.com/search?o=desc&q={}&s=stars',
        '!yt': 'https://www.youtube.com/results?search_query={}',
        '!g': 'https://www.google.com/search?q={}',
        }

c.completion.open_categories = ['searchengines', 'quickmarks', 'bookmarks', 'history', 'filesystem']
config.load_autoconfig(False) # load settings done via the gui
c.auto_save.session = True # save tabs on quit/restart

#colors

# keybinding changes
config.bind('=', 'cmd-set-text -s :open')
config.bind('cs', 'cmd-set-text -s :config-source')
config.bind('tH', 'config-cycle tabs.show multiple never')
config.bind('sH', 'config-cycle statusbar.show always never')
config.bind('T', 'hint links tab')
config.bind('pP', 'open -- {primary}')
config.bind('pp', 'open -- {clipboard}')
config.bind('pt', 'open -t -- {clipboard}')
config.bind('qm', 'macro-record')
config.bind('<ctrl-y>', 'spawn --userscript ytdl.sh')
config.bind('tT', 'config-cycle tabs.position top left')
config.bind('J', 'tab-prev')
config.bind('K', 'tab-next')
config.bind('gJ', 'tab-move +')
config.bind('gK', 'tab-move -')
config.bind('gm', 'tab-move')
config.bind('xx', 'config-cycle tabs.show multiple never')



# fonts
c.fonts.default_family = ['MonoLisa', 'GohuFont', 'monospace']
c.fonts.default_size = '13pt'
c.fonts.web.family.fixed = 'GohuFont Nerd Font'
c.fonts.web.family.sans_serif = 'GohuFont Nerd Font'
c.fonts.web.family.serif = 'GohuFont Nerd Font'
c.fonts.web.family.standard = 'GohuFont Nerd Font'

c.content.blocking.enabled = True
