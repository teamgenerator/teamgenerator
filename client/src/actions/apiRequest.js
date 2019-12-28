import {
  camelCase
} from 'change-case';

import pluralize from 'pluralize';

const makeRequestApiActionThread = (
  method,
  url,
  payload,
  actionType,
  typeExpected,
) => dispatch => {
  dispatch({
    type: `API_REQUEST ${url}`,
    // token: localStorage.getItem(localStorageKey),
    payload: payload === null || payload === undefined ? undefined : payload,
    actionType,
    method,
    url,
  });

  return fetch(`/api/v1${url}`, {
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
      // token: localStorage.getItem(localStorageKey),
    },
    method,
    body:
      payload === null || payload === undefined
        ? undefined
        : JSON.stringify(payload),
    // eslint-disable-next-line consistent-return
  }).then(response => {
    return response.text().then(rawBody => {

      dispatch({
        type: `API_RECEIVE ${url}`,
        rawBody,
      });

      if (response.status === 401 && url !== '/login') {
        dispatch({
          type: `LOG_OUT_FROM_EXPIRED_TOKEN`,
        });
      }

      let data = rawBody;
      try {
        data = JSON.parse(rawBody);
      } catch (e) {
        // keep the body raw
      }

      // if (response.status < 200 || response.status >= 300) {
      //   let msg = '';
      //   if (
      //     response.status === 500 ||
      //     response.headers.get('content-type').includes('text/html')
      //   ) {
      //     msg = 'Internal error. Please contact support to resolve this issue.';
      //   } else {
      //     msg = body.message;
      //   }

      //   dispatch({
      //     type: 'ALERT_ERROR',
      //     message: msg,
      //   });
      //   return Promise.reject(response.status);
      // }

      // if (response.headers.get('token')) {
      //   localStorage.setItem(localStorageKey, response.headers.get('token'));
      // }

      // if (url === '/login') {
      //   // build the body of the reducer.
      //   const loginState = {
      //     ...body.data[0].attributes,
      //     j_permissions: JSON.parse(body.data[0].attributes.j_permissions),
      //   };

      //   // persist it to the localStorage with the given key.
      //   // localStorage.setItem(
      //   //   'jabc-app-login-state',
      //   //   JSON.stringify(loginState)
      //   // );

      //   dispatch({
      //     type: 'REPLACE_LOGIN',
      //     payload: loginState,
      //   });
      // }


      if (data && Array.isArray(data)) {

        const camelCasedData = data.map(datum => {
          const obj = Object.entries(datum).reduce((newObj, [oldKey, val]) => {
            newObj[camelCase(oldKey)] = val;
            return newObj;
          }, {});

          obj.type = typeExpected;
          return obj;
        });

        const entityMap = camelCasedData.reduce((entityMap, datum) => {
          entityMap[datum.id] = datum;
          return entityMap;
        }, {});


        dispatch({
          type: `${pluralize(typeExpected).toUpperCase()}_${actionType.toUpperCase()}`,
          data: entityMap,
        });
      }

      // if (body.data && Array.isArray(body.data)) {
      //   const types = body.data.reduce((acc, entity) => {
      //     if (!acc.includes(entity.type)) {
      //       acc.push(entity.type);
      //     }

      //     return acc;
      //   }, []);

      //   types.forEach(type => {
      //     const entities = body.data.filter(entity => entity.type === type);

      //     const entityMap = entities.reduce((acc, entity) => {
      //       return {
      //         ...acc,
      //         [entity.id]: entity.attributes,
      //       };
      //     }, {});

      //     dispatch({
      //       type: `${actionType.toUpperCase()}_${pluralize(
      //         type
      //       ).toUpperCase()}`,
      //       payload: entityMap,
      //     });
      //   });
      // }

      return Promise.resolve(data);
    });
  });
};

export default makeRequestApiActionThread;
