import { useEffect } from 'react'
import api from '../utils/api'

export function Home() {
  useEffect(() => {
    api.get('/')
  }, [])
  return <>Hey</>
}
