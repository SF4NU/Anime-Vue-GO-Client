export function compareLengths(strings) {
  const checkStrings = strings.filter((str) => str.trim() !== "");

  if (checkStrings.length === 0) {
    return;
  }

  checkStrings.sort((a, b) => a.length - b.length);

  return checkStrings[0];
}
