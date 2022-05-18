db.createUser({
  user: 'ecomm',
  pwd: 'password',
  roles: [{
    role: 'readWrite',
    db: 'Ecommerce'
  }]
})

