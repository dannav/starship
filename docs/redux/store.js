import { createStore, applyMiddleware } from 'redux'
import { composeWithDevTools } from 'redux-devtools-extension'
import thunkMiddleware from 'redux-thunk'
import fetch from 'isomorphic-unfetch'

const initialState = {
  content: [],
}

export const actionTypes = {
  GET_PAGE: 'GET_PAGE'
}

// REDUCERS
export const reducer = (state = initialState, action) => {
  switch (action.type) {
    case actionTypes.GET_PAGE:
      return {
        ...state,
        content: action.content,
      }
    default: return state
  }
}

// ACTIONS
export const getContent = (isServer, sRes, folder, page) => async (dispatch, getState) => {
  const res = await fetch(`http://localhost:6060/static/${folder}/${page}.md`)

  if (res.status == 404) {
    if (isServer) {
      sRes.writeHead(302, {
        Location: 'http://localhost:6060/not-found'
      })
      sRes.end()
      return
    }

    window.location = '/not-found'
    return
  }

  const content = await res.text()
  return dispatch({type: actionTypes.GET_PAGE, content})
}

export const initStore = (initialState = initialState) => {
  return createStore(
    reducer,
    initialState,
    composeWithDevTools(applyMiddleware(thunkMiddleware))
  )
}