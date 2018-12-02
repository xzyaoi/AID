function getMenus () {
  return [
        { 'header': 'Admin' },
        { 'href': '/', 'title': 'Home', 'icon': 'home' },
        { 'href': '/local/package', 'title': 'Packages', 'icon': 'view_list' },
        { 'href': '/local/pretrained', 'title': 'Pre-trained', 'icon': 'view_list' },
        { 'header': 'System' },
        { 'href': '/settings', 'title': 'Settings', 'icon': 'settings' },
        { 'href': '/login', 'icon': 'lock', 'title': 'Logout' }
  ]
}

export {
    getMenus
}