import axios from 'axios'

const api = axios.create({ baseURL: process.env.REACT_APP_BACKEND_URL })

api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response.status === 401) {
      localStorage.removeItem('name')
      localStorage.removeItem('avatar')
      localStorage.removeItem('email')
      localStorage.removeItem('user_id')
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

export default api
