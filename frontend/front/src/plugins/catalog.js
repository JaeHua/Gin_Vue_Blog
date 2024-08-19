export function offsetDomTop (el) {
  let top = 0
  while (el) {
    top += el.offsetTop
    el = el.offsetParent
  }
  return top
}

export function scrollToToc (toc) {
  const element = document.getElementById(toc.id)
  if (element) {
    window.scrollTo({
      top: element.offsetTop - 20,
      behavior: 'smooth'
    })
  }
}
