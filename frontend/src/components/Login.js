import React, { useEffect } from 'react'
import { useLocation, useNavigate } from 'react-router-dom'
import axios from 'axios'

const backendUrl = process.env.REACT_APP_BACKEND_URL

export function Login() {
  const handleLogin = () => {
    window.location.href = backendUrl + '/auth/google/login'
  }
  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="text-center">
        <button
          className="inline-flex items-center px-6 py-3 font-semibold text-gray-700 bg-white border border-gray-300 rounded-full shadow-md hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500"
          onClick={handleLogin}
        >
          <img src="google-logo.jpg" className="w-6 h-6 mr-3" />
          Sign in with Google
        </button>
        <pre id="user-info" className="mt-4 text-left"></pre>
      </div>
    </div>
  )
}

export function LoginCallBack() {
  const params = new URLSearchParams(useLocation().search)
  const code = params.get('code')
  const navigate = useNavigate()
  useEffect(() => {
    axios
      .get(backendUrl + '/auth/google/callback' + `?code=${code}`)
      .then((res) => {
        if (res.status === 200) {
          console.log(res.data)
          localStorage.setItem('name', res.data.user.Name)
          localStorage.setItem('avatar', res.data.user.avatar)
          localStorage.setItem('email', res.data.user.Email)
          localStorage.setItem('user_id', res.data.user.UserID)
          setTimeout(() => {
            navigate('/')
          }, 5000)
        } else {
          navigate('/login')
        }
      })
  }, [code])
  return (
    <div className="bg-gradient-to-r from-white-500 to-indigo-600 flex items-center justify-center min-h-screen">
      <div className="text-center">
        <div className="loader rounded-full h-24 w-24 border-4 border-white border-opacity-25 mx-auto mb-6"></div>
        <h1 className="text-3xl font-bold text-black mb-2">
          Loading, Your will be redirect in 59 minutes...
        </h1>
        <p className="text-lg text-blank opacity-80">
          I'm kidding, you will be redirect shortly.
        </p>
      </div>
    </div>
  )
}
