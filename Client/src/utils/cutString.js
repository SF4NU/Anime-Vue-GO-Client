export function cutString(str) {
  if (str.length > 30) {
    return str.substring(0, 30) + "...";
  }
  return str;
}
