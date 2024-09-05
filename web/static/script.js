const $form = document.querySelector('form')
const $fileInput = document.querySelector('#files')
const $progress = document.querySelector('#progress-bar')
const $status = document.querySelector('#status')

$form.addEventListener('submit', e => {
    e.preventDefault()

    const formData = new FormData($form)

    const xhr = new XMLHttpRequest()
    xhr.open('POST', 'http://127.0.0.1:5000/upload', true)

    xhr.upload.addEventListener('progress', e => {
        const percent = (e.loaded / e.total) * 100
        $progress.value = percent
    })

    xhr.upload.addEventListener('load', () => {
        $progress.value = 100
    })

    xhr.upload.addEventListener('error', () => {
        $progress.value = 0
    })

    xhr.addEventListener('load', () => {
        if (xhr.status === 200) {
            $status.textContent = 'Upload complete'
            $fileInput.value = ''
            $progress.value = 0

            // updateRecentUploads()
        } else {
            $status.textContent = 'Upload failed'
        }
    })

    xhr.send(formData)
})
