// Create Folder popup
const createFolder = document.getElementById('createFolder')
const modal = document.getElementById('modal')
const closeModal = document.getElementById('closeModal')

createFolder.addEventListener('click', () => {
  modal.classList.remove('hidden')
})

closeModal.addEventListener('click', () => {
  modal.classList.add('hidden')
})

window.addEventListener('click', (e) => {
  if (e.target == modal) {
    modal.classList.add('hidden')
  }
})

document
  .getElementById('create-folder')
  .addEventListener('htmx:afterRequest', function (event) {
    res = event.detail.xhr
    data = JSON.parse(res.response)
    if (res.status === 200) {
      modal.classList.add('hidden')
      window.location.href = `/folder/${data.data.ID}`
    } else {
      document.getElementById('error-message').textContent = data.message
    }
  })

// Upload folder / file popup
