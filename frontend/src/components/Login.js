import React from 'react'

function Login() {
  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="text-center">
        <button className="inline-flex items-center px-6 py-3 font-semibold text-gray-700 bg-white border border-gray-300 rounded-full shadow-md hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500">
          <img src="google-logo.jpg" className="w-6 h-6 mr-3" />
          Sign in with Google
        </button>
        <pre id="user-info" className="mt-4 text-left"></pre>
      </div>
    </div>
  )
}

export default Login
