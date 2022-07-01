import React from 'react'

export default function ArrayComponent() {
  return (
    <div class="flex flex-col overflow-hidden items-center h-screen p-60 w-screen mb-30">
    <div>
      <div class="text-7xl">"asd"</div>
     </div>
  <div class="form-control w-full max-w-xs mt-48">
    <span>
        <label class="block text-gray-700 text-sm font-bold mb-2" for="inline-full-name">
        </label>
    </span>
  <input type="text" placeholder="Enter your string" class="input input-bordered w-full max-w-xs " />
  <div class="divider"></div>
  <div class="grid grid-cols-3 pb-5 gap-2">
    <button class="btn btn-xs sm:btn-sm md:btn-md lg:btn-sm">Palindrome</button>
    <button class="btn btn-xs sm:btn-sm md:btn-md lg:btn-sm">RemoveDups</button>
    <button class="btn btn-xs sm:btn-sm md:btn-md lg:btn-sm">toLowercase</button>
  </div>
</div>
    </div>
  
    )
}
