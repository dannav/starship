import { createStore, applyMiddleware } from 'redux'
import { composeWithDevTools } from 'redux-devtools-extension'
import thunkMiddleware from 'redux-thunk'
import fetch from 'isomorphic-unfetch'

const initialState = {
  posts: [],
}

export const actionTypes = {
  GET_POSTS: 'GET_POSTS'
}

// REDUCERS
export const reducer = (state = initialState, action) => {
  switch (action.type) {
    case actionTypes.GET_POSTS:
      return {
        ...state,
        posts: action.results,
      }
    default: return state
  }
}

// ACTIONS
export const getPosts = isServer => async (dispatch, getState) => {
  let { posts } = getState()
  if (!posts.length) {
    const res = await fetch('http://localhost:6060/static/posts.json')
    posts = await res.json()

    return dispatch({type: actionTypes.GET_POSTS, results: posts.results})
  }

  return
}

export const initStore = (initialState = initialState) => {
  return createStore(
    reducer,
    initialState,
    composeWithDevTools(applyMiddleware(thunkMiddleware))
  )
}