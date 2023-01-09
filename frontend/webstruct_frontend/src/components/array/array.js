import React from 'react'
import axios from 'axios'
import arrayHelper from '../../helpers/arrayop'
import { motion } from 'framer-motion'
import toast, { Toaster } from 'react-hot-toast'

export default function ArrayComponent() {
  // get url from config

  //dotenv
  const url = process.env.REACT_APP_API_URL

  const [array, setArray] = React.useState([])
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
      .post(url + '/structures/operation', { Entity: 'string', Type: 'get' })
      .then((response) => {
        setArray(response.data.result)
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
          setArray(value)
        })
        .catch((error) => {
          console.log(error)
        })
    }
  }

  function handleRemoveDups() {
    arrayHelper
      .removeDups()
      .then((response) => {
        setArray(response.data.result)
      })
      .catch((error) => {
        console.log(error)
      })
  }

  function handleToLowcase() {
    arrayHelper
      .lowCase()
      .then((response) => {
        setArray(response.data.result)
      })
      .catch((error) => {
        console.log(error)
      })
  }

  function handlePalindrome() {
    arrayHelper
      .isPalindrome()
      .then((response) => {
        if (response.data.message === 'true') {
          toast('This is a palindrome!', {
            type: 'success',
            position: 'top-right',
            autoClose: 4000,
            hideProgressBar: false,
            closeOnClick: true,
            pauseOnHover: true,
            style: {
              backgroundColor: '#00b894',
              color: '#fff',
              fontSize: '1.5rem',
              fontWeight: 'bold',
              borderRadius: '12px',
            },
          })
        } else {
          toast('This is not a palindrome!', {
            type: 'error',
            position: 'top-right',
            autoClose: 4000,
            hideProgressBar: false,
            closeOnClick: true,
            pauseOnHover: true,
            style: {
              backgroundColor: '#b00d64',
              color: '#fff',
              fontSize: '1.5rem',
              fontWeight: 'bold',
              borderRadius: '12px',
            },
          })
        }
      })
      .catch((error) => {
        console.log(error)
      })
  }

  return (
    <div className="flex flex-col overflow-hidden items-center p-60 mb-30">
      {array.length === 0 ? (
        <motion.div
          key={array}
          variants={variants}
          initial="hidden"
          animate={'visible'}
        >
          <div className="text-4xl font-bold">empty</div>
        </motion.div>
      ) : (
        <motion.div
          key={array}
          variants={variants}
          initial="hidden"
          animate={'visible'}
        >
          <div className="text-4xl font-bold">{array}</div>
        </motion.div>
      )}
      <div className="form-control w-full max-w-xs mt-48">
        <span>
          <label
            className="block text-gray-700 text-sm font-bold mb-2"
            htmlFor="inline-full-name"
          ></label>
        </span>

        <input
          maxLength={32}
          onKeyPress={onEnterKeyPress}
          type="text"
          placeholder="Enter your string"
          className="input input-bordered w-full max-w-xs "
        />
        <div className="divider"></div>
        <div className="grid grid-cols-3 pb-5 gap-2">
          <button
            className="btn btn-xs sm:btn-sm md:btn-md lg:btn-sm"
            onClick={handlePalindrome}
          >
            Palindrome
          </button>
          <button
            className="btn btn-xs sm:btn-sm md:btn-md lg:btn-sm"
            onClick={handleRemoveDups}
          >
            RemoveDups
          </button>
          <button
            className="btn btn-xs sm:btn-sm md:btn-md lg:btn-sm"
            onClick={handleToLowcase}
          >
            toLowercase
          </button>
        </div>
      </div>
    </div>
  )
}
