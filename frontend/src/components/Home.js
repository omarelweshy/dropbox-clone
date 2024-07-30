import { useEffect } from 'react'
import api from '../utils/api'

export function Home() {
  useEffect(() => {
    api.get('/').then((res) => res.data)
  }, [])
  return (
    <>
      <div class="container mx-auto p-6">
        <div class="bg-white p-4 rounded-lg">
          <div class="mb-4">
            <input
              type="text"
              placeholder="Search"
              class="w-full bg-gray-200 text-gray-700 px-4 py-2 rounded-lg focus:outline-none"
            />
          </div>

          <div class="flex flex-wrap items-center space-x-4">
            <button class="bg-white text-gray-700 px-4 py-2 rounded-lg shadow-sm border border-gray-300 flex items-center">
              <svg
                class="h-5 w-5 text-gray-700 mr-2"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M12 4v16m8-8H4"
                />
              </svg>
              Create Folder
            </button>
            <button class="bg-white text-gray-700 px-4 py-2 rounded-lg shadow-sm border border-gray-300 flex items-center">
              <svg
                class="h-5 w-5 text-gray-700 mr-2"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M3 16.5V19a2.5 2.5 0 002.5 2.5h13A2.5 2.5 0 0021 19v-2.5m-4-4l-5-5m0 0l-5 5m5-5V19"
                />
              </svg>
              Upload file or folder
            </button>
          </div>
        </div>
      </div>

      <div class="container mx-auto p-6">
        <h1 class="text-2xl font-semibold mb-4">Files and Folders</h1>
        <div class="overflow-x-auto bg-white shadow-md rounded-lg">
          <table class="min-w-full bg-white">
            <thead class="bg-gray-200 text-gray-600">
              <tr>
                <th class="py-3 px-4 text-left text-sm font-semibold">Name</th>
                <th class="py-3 px-4 text-left text-sm font-semibold">Type</th>
                <th class="py-3 px-4 text-left text-sm font-semibold">Size</th>
                <th class="py-3 px-4 text-left text-sm font-semibold">
                  Last Modified
                </th>
                <th class="py-3 px-4 text-left text-sm font-semibold">
                  Actions
                </th>
              </tr>
            </thead>
            <tbody class="text-gray-700">
              <tr class="border-b border-gray-200 hover:bg-gray-50">
                <td class="py-3 px-4 flex items-center">
                  <svg
                    class="h-5 w-5 text-blue-500 mr-2"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M3 7h18M3 7a2 2 0 012-2h14a2 2 0 012 2v10a2 2 0 01-2 2H5a2 2 0 01-2-2V7z"
                    />
                  </svg>
                  <span>Documents</span>
                </td>
                <td class="py-3 px-4">Folder</td>
                <td class="py-3 px-4">-</td>
                <td class="py-3 px-4">2024-07-30</td>
                <td class="py-3 px-4">
                  <button class="text-blue-600 hover:underline">Open</button>
                </td>
              </tr>
              <tr class="border-b border-gray-200 hover:bg-gray-50">
                <td class="py-3 px-4 flex items-center">
                  <svg
                    class="h-5 w-5 text-green-500 mr-2"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M12 12v4m0 0v4m0-4h4m-4 0H8m8-8v4H8V8a2 2 0 012-2h4a2 2 0 012 2z"
                    />
                  </svg>
                  <span>example.pdf</span>
                </td>
                <td class="py-3 px-4">File</td>
                <td class="py-3 px-4">1.2 MB</td>
                <td class="py-3 px-4">2024-07-30</td>
                <td class="py-3 px-4">
                  <button class="text-blue-600 hover:underline">
                    Download
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </>
  )
}

// <div class="container mx-auto p-4">
//   <div class="bg-white-900 p-8 rounded-lg border border-gray-700">
//     <div class="flex flex-col items-center justify-center border-black-600 p-16">
//       <svg
//         class="h-16 w-16 mb-4"
//         fill="none"
//         stroke="currentColor"
//         viewBox="0 0 24 24"
//         xmlns="http://www.w3.org/2000/svg"
//       >
//         <path
//           stroke-linecap="round"
//           stroke-linejoin="round"
//           stroke-width="2"
//           d="M3 16.5V19a2.5 2.5 0 002.5 2.5h13A2.5 2.5 0 0021 19v-2.5m-4-4l-5-5m0 0l-5 5m5-5V19"
//         />
//       </svg>
//       <p class="text-lg mb-2">Drop files here to upload</p>
//       <button class="bg-gray-700 text-white px-4 py-2 rounded-lg flex items-center">
//         Upload
//       </button>
//     </div>
//   </div>
// </div>
//
