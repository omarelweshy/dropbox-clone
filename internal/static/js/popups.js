const createFolder = document.getElementById('createFolder')
const modal = document.getElementById('modal')
const closeModal = document.getElementById('closeModal')

createFolder.addEventListener('click', () => {
  console.log('hery')
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

// document
//   .getElementById('create-folder-form')
//   .addEventListener('htmx:configRequest', function (evt) {
//     // Set the Content-Type header to JSON
//     evt.detail.headers['Content-Type'] = 'application/json'
//
//     // Serialize form data into JSON
//     const form = evt.target
//     const formData = new FormData(form)
//     const jsonObject = {}
//     formData.forEach((value, key) => {
//       jsonObject[key] = value
//     })
//
//     // Replace the body with the JSON string
//     evt.detail.body = JSON.stringify(jsonObject)
//   })
