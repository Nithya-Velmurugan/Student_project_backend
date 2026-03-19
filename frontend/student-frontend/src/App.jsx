import React, { useState } from 'react';
import './App.css';

function App() {
  const [isLoginView, setIsLoginView] = useState(true);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [successMsg, setSuccessMsg] = useState('');
  const [loading, setLoading] = useState(false);
  const [token, setToken] = useState(localStorage.getItem('token') || '');

  const toggleView = () => {
    setIsLoginView(!isLoginView);
    setError('');
    setSuccessMsg('');
    setEmail('');
    setPassword('');
  };

  const handleAuth = async (e) => {
    e.preventDefault();
    setError('');
    setSuccessMsg('');
    setLoading(true);

    const endpoint = isLoginView ? '/auth/login' : '/auth/register';
    
    try {
      const response = await fetch(`http://localhost:8082${endpoint}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password, role: isLoginView ? undefined : "Student" }),
      });

      const data = await response.json();

      if (!response.ok) {
        throw new Error(data.error || 'Authentication failed');
      }

      if (isLoginView) {
        setToken(data.token);
        localStorage.setItem('token', data.token);
      } else {
        setSuccessMsg('Registration successful! Please sign in.');
        setIsLoginView(true);
        setPassword('');
      }
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const handleLogout = () => {
    setToken('');
    localStorage.removeItem('token');
  };

  if (token) {
    return (
      <div className="login-container" style={{ flexDirection: 'column', gap: '1rem' }}>
        <div className="login-card">
          <div className="icon-container" style={{ backgroundColor: '#10b981' }}>
            <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
              <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path>
              <polyline points="22 4 12 14.01 9 11.01"></polyline>
            </svg>
          </div>
          <h1>Authentication Successful!</h1>
          <p className="subtitle">You have successfully authenticated to the system.</p>
          <button onClick={handleLogout} className="btn-primary" style={{ marginTop: '1rem' }}>
            Logout
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className="login-container">
      <div className="login-card">
        <div className="icon-container">
          {/* Change icon based on view (grad cap for login, user plus for register) */}
          <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
            <path d={isLoginView ? "M22 10v6M2 10l10-5 10 5-10 5z" : "M16 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"} />
            <path d={isLoginView ? "M6 12v5c3 3 9 3 12 0v-5" : "M8.5 3a4 4 0 1 0 0 8 4 4 0 0 0 0-8z"} />
            {!isLoginView && <line x1="20" y1="8" x2="20" y2="14" />}
            {!isLoginView && <line x1="23" y1="11" x2="17" y2="11" />}
          </svg>
        </div>
        <h1>{isLoginView ? 'Welcome Back' : 'Create Account'}</h1>
        <p className="subtitle">Student Management System</p>
        
        {error && <div style={{ color: '#ef4444', marginBottom: '1rem', fontSize: '0.875rem', fontWeight: '500' }}>{error}</div>}
        {successMsg && <div style={{ color: '#10b981', marginBottom: '1rem', fontSize: '0.875rem', fontWeight: '500' }}>{successMsg}</div>}

        <form onSubmit={handleAuth}>
          <div className="form-group">
            <label htmlFor="email">Email</label>
            <input 
              type="email" 
              id="email" 
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              placeholder="your.email@school.com" 
              required 
            />
          </div>
          
          <div className="form-group">
            <label htmlFor="password">Password</label>
            <input 
              type="password" 
              id="password" 
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              placeholder="••••••••" 
              required 
              minLength={6}
            />
          </div>
          
          {isLoginView && (
            <div className="form-actions">
              <label className="remember-me">
                <input type="checkbox" />
                <span>Remember me</span>
              </label>
              <a href="#" className="forgot-password">Forgot password?</a>
            </div>
          )}
          
          <button type="submit" className="btn-primary" disabled={loading} style={{ marginTop: isLoginView ? '0' : '1.5rem' }}>
            {loading ? 'Processing...' : (isLoginView ? 'Sign In' : 'Sign Up')}
          </button>
        </form>
        
        <p className="footer-text">
          {isLoginView ? "Don't have an account? " : "Already have an account? "}
          <span className="footer-link" style={{cursor: 'pointer'}} onClick={toggleView}>
            {isLoginView ? 'Sign Up' : 'Sign In'}
          </span>
        </p>
      </div>
    </div>
  );
}

export default App;
