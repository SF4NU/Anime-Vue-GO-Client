export const timeout = (ms) => {
  return new Promise((res, rej) => {
    setTimeout(() => {
      res();
    }, ms);
  });
};
