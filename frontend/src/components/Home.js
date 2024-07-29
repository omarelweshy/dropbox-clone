import { useEffect } from 'react'
import api from '../utils/api'

export function Home() {
  useEffect(() => {
    api.get('/').then((res) => res.data)
  }, [])
  return <>Hey</>
}
