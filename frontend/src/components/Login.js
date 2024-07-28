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
          navigate('/')
        } else {
          navigate('/login')
        }
      })
  }, [code])
}
