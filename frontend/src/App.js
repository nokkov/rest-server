import './App.css';

function App() {
  return (
    <div className="App">
      <div className="container">
        <div className="logo">
          <img src="ULTIMATE-URL-SHORTENER.gif" alt="logo" />
        </div>
      </div>
      <div className="input-fields">
        <form id="urlForm">
          <input type="text" id="longUrlField" required></input>
          <input type="text" id="shortUrlField" readOnly></input>
          <button type="button" onClick="submitForm()">SEND NUDES!</button>
        </form>  
      </div>
      <script defer src="./handler.js"></script>
    </div>
  );
}

export default App;
