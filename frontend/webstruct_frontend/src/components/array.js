import React from 'react'
import axios from 'axios'
import arrayHelper from '../helpers/arrayop'
import { motion } from 'framer-motion'

export default function ArrayComponent() {
  // get url from config

  //dotenv
  const url = process.env.REACT_APP_API_URL

  console.log(url)
  const [array, setArray] = React.useState([])

  const variants = {
    visible: {
      opacity: 1,
      y: 0,
      transition: {
        duration: 1,
        ease: 'easeOut',
        duration: 0.4 
      }
    },
    hidden: {
      opacity: 0,
      y: -20
    }
  }
  const initial = {
    opacity: 0,
    scale : 0.5,
    transition: {
      duration: 1,
      ease: [0.6, 0.05, -0.01, 0.9]
    }
  }


  React.useEffect(() => {
    axios
      .post(url + 'arrays?op=get')
      .then((response) => {
        setArray(response.data.array)
      })
      .catch((error) => {
        console.log(error)
      })
  }, [])

  const onEnterKeyPress = (e) => {
    if (e.key === 'Enter') {
      const value = e.target.value
      arrayHelper
        .addArray(value)
        .then((response) => {
          setArray(response.data.array)
        })
        .catch((error) => {
          console.log(error)
        })
      setArray(value)
    }
  }

  function handleRemoveDups() {
      arrayHelper
        .removeDups()
        .then((response) => {
          setArray(response.data.array)
        })
        .catch((error) => {
          console.log(error)
        })
    }
  

  function handleToLowcase() {
      arrayHelper
        .lowCase()
        .then((response) => {
          setArray(response.data.array)
        })
        .catch((error) => {
          console.log(error)
        })
    }

    return (
      <div class="flex flex-col overflow-hidden items-center h-screen p-60 w-screen mb-30">
        <div>
          {array.length === 0 ? 
            <motion.div key={array} variants={variants} initial="hidden" animate = {"visible"}>
            <div class="text-4xl font-bold">empty</div>
            </motion.div>
           : 
            <motion.div key={array} variants={variants} initial="hidden" animate = {"visible"}>
            <div class="text-4xl font-bold">{array}</div>
            </motion.div>
          }
        </div>

        <div class="form-control w-full max-w-xs mt-48">
          <span>
            <label
              class="block text-gray-700 text-sm font-bold mb-2"
              for="inline-full-name"
            ></label>
          </span>

          <input
            maxLength={32}
            onKeyPress={onEnterKeyPress}
            type="text"
            placeholder="Enter your string"
            class="input input-bordered w-full max-w-xs "
          />
          <div class="divider"></div>
          <div class="grid grid-cols-3 pb-5 gap-2">
            <button class="btn btn-xs sm:btn-sm md:btn-md lg:btn-sm">
              Palindrome
            </button>
            <button class="btn btn-xs sm:btn-sm md:btn-md lg:btn-sm" onClick={handleRemoveDups}>
              RemoveDups
            </button>
            <button class="btn btn-xs sm:btn-sm md:btn-md lg:btn-sm" onClick={handleToLowcase}>
              toLowercase
            </button>
          </div>
        </div>
      </div>
    )
  }

