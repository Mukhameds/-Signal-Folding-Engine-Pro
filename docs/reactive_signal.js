function activateReactive() {
  if (!currentComp) {
    console.warn("No structure loaded");
    return;
  }

  console.log("🧠 Reactive Mode ON: starting φ/ψ signal propagation...");

  const interval = setInterval(() => {
    const phi = (Math.random() * 360 - 180).toFixed(2);
    const psi = (Math.random() * 360 - 180).toFixed(2);
    const message = `Signal → φ=${phi}°, ψ=${psi}° → Torsion Adjusted`;
    console.log(message);

    const readout = document.getElementById("signal-readout");
    if (readout) {
      readout.textContent = `Signal φ: ${phi}°, ψ: ${psi}°`;
    }
  }, 1000);

  setTimeout(() => {
    clearInterval(interval);
    console.log("🛑 Reactive Mode paused");
  }, 10000);
}
