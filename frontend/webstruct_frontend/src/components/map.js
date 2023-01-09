import React from 'react'
import axios from 'axios'
import HashMapHelper from '../helpers/hashop.js'
import { motion } from 'framer-motion'
import toast, { Toaster } from 'react-hot-toast'
import FlowInstance from './flow/flow.js'
import hashMapHelper from '../helpers/hashop.js'

export default function MapComponent() {
  const url = process.env.REACT_APP_API_URL
  const keyInput = React.createRef()
  const valueInput = React.createRef()

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
      .post(url + '/structures/operation', { Entity: 'hashmap', Type: 'get' })
      .then((response) => {
        setMap(response.data.result.entries)
      })
      .catch((error) => {
        console.log(error)
      })
  }, [])

  const onEnterKeyPress = (e) => {
    if (e.key === 'Enter') {
      const key = keyInput.current.value
      const value = valueInput.current.value
      if (keyInput.current.value !== '' && valueInput.current.value !== '') {
        hashMapHelper
          .setKV(key, value)
          .then(
            (response) => {
              hashMapHelper.getMap(response.data.result)
              setMap(response.data.result)
            },
            (error) => {
              console.log(error)
            }
          )
          .catch((error) => {
            console.log(error)
          })
      } else {
        toast.error('Please enter a key and value')
      }
    }
  }

  const createMap = () => {
    hashMapHelper.addMap()
  }

  const ClearMap = () => {
    hashMapHelper.deleteMap()
    setMap(hashMapHelper.getMap())
  }

  return (
    <div className="flex flex-col overflow-hidden items-center p-60 mb-30">
      <div className="h-44 overflow-y-auto w-72 overflow-x-hidden  block scroll-smooth">
        {Object.keys(map).length === 0 ? (
          <motion.div variants={variants} initial="hidden" animate="visible">
            <div className="grid grid-cols-2 gap-4 rounded-lg text-gray-700 text-sm font-bold text-center">
              <div className="px-2 py-2  text-slate-500 rounded-lg">Key</div>
              <div className="px-2 py-2  text-slate-500 rounded-lg">Value</div>
            </div>
          </motion.div>
        ) : (
          <div className="transition-opacity transition-duration-500">
            <motion.div variants={variants} initial="hidden" animate="visible">
              <div className="grid grid-cols-2 gap-4 rounded-lg text-gray-700 text-sm font-bold text-center">
                <div className="px-2 py-2  text-slate-500 rounded-lg">Key</div>
                <div className="px-4 py-2 text-slate-500 rounded-lg">Value</div>
              </div>
            </motion.div>
            <motion.div
              className="grid grid-cols-2 gap-4 rounded-lg text-center"
              variants={variants}
              initial="hidden"
              animate="visible"
            >
              {Object.entries(map).map(([key, value]) => (
                <React.Fragment key={key}>
                  <motion.div
                    variants={variants}
                    initial="hidden"
                    animate="visible"
                    className="px-2 bg-gray-600 rounded-xl"
                  >
                    <strong className="block text-gray-200 text-sm font-bold mb-2">
                      {key}
                    </strong>
                  </motion.div>
                  <motion.div
                    variants={variants}
                    initial="hidden"
                    animate="visible"
                    className="px-2  bg-gray-600 rounded-xl "
                  >
                    <span className="block text-gray-200 text-xs">{value}</span>
                  </motion.div>
                </React.Fragment>
              ))}
            </motion.div>
          </div>
        )}
      </div>
      <div className="form-control w-full max-w-xs mt-4">
        <span>
          <label
            className="block text-gray-700 text-sm font-bold mb-2"
            htmlFor="inline-full-name"
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
            ref={valueInput}
          />
        </div>
        <div className="divider"></div>
        <div className="grid grid-cols-3 pb-5 gap-2">
          <button
            className="btn btn-xs sm:btn-sm md:btn-md lg:btn-sm"
            onClick={createMap}
          >
            Create Map
          </button>
          <button
            className="btn btn-xs sm:btn-sm md:btn-md lg:btn-sm"
            onClick={ClearMap}
          >
            ClearMap
          </button>
          <button className="btn btn-xs sm:btn-sm md:btn-md lg:btn-sm">
            Op3
          </button>
        </div>
      </div>
    </div>
  )
}
