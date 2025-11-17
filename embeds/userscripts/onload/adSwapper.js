/**
 * Swaps out the content of an ad container with a Navlink Ad.
 * https://dimden.dev/navlinkads/
 */

document.querySelectorAll('.ad-container').forEach((elem) => {
  elem.innerHTML = `<iframe width="90" height="90" style="border:none" src="https://dimden.neocities.org/navlink/" name="neolink"></iframe>`
})
