import React from 'react'
import axios from 'axios'
import HashMapHelper from '../../helpers/hashop.js'
import { motion } from 'framer-motion'
import toast, { Toaster } from 'react-hot-toast'
import FlowInstance from '../flow/flow.js'
import hashMapHelper from '../../helpers/hashop.js'

export default function MapComponent() {
  const url = process.env.REACT_APP_API_URL
  const keyInput = React.createRef()
  const valueInput = React.createRef()

  console.log(url)
  const [map, setMap] = React.useState([])

  const variants = {
    visible: {
      opacity: 1,
      y: 0,
      transition: {
        duration: 1,
        ease: 'easeOut',
        duration: 0.4,
      },
    },
    hidden: {
      opacity: 0,
      y: -20,
    },
  }
  const initial = {
    opacity: 0,
    scale: 0.5,
    transition: {
      duration: 1,
      ease: [0.6, 0.05, -0.01, 0.9],
    },
  }

  React.useEffect(() => {
    axios
      .post(url + 'structures/operation', { Entity: 'HashMap', Type: 'get' })
      .then((response) => {
        setMap(response.result)
      })
      .catch((error) => {
        console.log(error)
      })
  }, [])

  const onEnterKeyPress = (e) => {
    if (e.key === 'Enter') {
      const value = e.target.value
      hashMapHelper
        .addMap(value)
        .then((response) => {
          setMap(response.data.array)
        })
        .catch((error) => {
          console.log(error)
        })
      setMap(value)
    }
  }

  return (
    <div className="flex flex-col overflow-hidden items-center h-screen p-60 w-screen mb-30">
      <div className="bg-red-800 w-4/5 h-screen border-4 rounded border-slate-700">
        <FlowInstance />
      </div>

      <div className="form-control w-full max-w-xs mt-48">
        <span>
          <label
            className="block text-gray-700 text-sm font-bold mb-2"
            for="inline-full-name"
          ></label>
        </span>
        {/** A horizontal tailwind stack */}
        <div className="flex flex-row space-x-3">
          <input
            maxLength={16}
            onKeyPress={onEnterKeyPress}
            type="text"
            placeholder="Enter your key"
            className="input input-bordered w-full max-w-xs"
            ref={keyInput}
          />
          <input
            maxLength={16}
            onKeyPress={onEnterKeyPress}
            type="text"
            placeholder="Enter your value"
            className="input input-bordered w-full max-w-xs "
            res={valueInput}
          />
        </div>
        <div className="divider"></div>
        <div className="grid grid-cols-3 pb-5 gap-2">
          <button className="btn btn-xs sm:btn-sm md:btn-md lg:btn-sm">
            Op1
          </button>
          <button className="btn btn-xs sm:btn-sm md:btn-md lg:btn-sm">
            Op2
          </button>
          <button className="btn btn-xs sm:btn-sm md:btn-md lg:btn-sm">
            Op3
          </button>
        </div>
      </div>
    </div>
  )
}
