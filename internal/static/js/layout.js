document.addEventListener('DOMContentLoaded', function () {
  const userMenuButton = document.getElementById('user-menu-button')
  const userMenu = document.getElementById('user-menu')

  userMenuButton.addEventListener('click', function (event) {
    event.stopPropagation()
    userMenu.classList.toggle('hidden')
  })

  document.addEventListener('click', function (event) {
    if (!userMenu.contains(event.target)) {
      userMenu.classList.add('hidden')
    }
  })
})
