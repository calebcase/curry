// This file is generated - do not edit.

package curry

func A2R1[A0, A1, R0 any](f func(a0 A0, a1 A1) R0) func(a0 A0) func(a1 A1) R0 {
	return func(a0 A0) func(a1 A1) R0 {
		return func(a1 A1) R0 {
			return f(a0, a1)
		}
	}
}

func A3R1[A0, A1, A2, R0 any](f func(a0 A0, a1 A1, a2 A2) R0) func(a0 A0) func(a1 A1) func(a2 A2) R0 {
	return func(a0 A0) func(a1 A1) func(a2 A2) R0 {
		return func(a1 A1) func(a2 A2) R0 {
			return func(a2 A2) R0 {
				return f(a0, a1, a2)
			}
		}
	}
}

func A4R1[A0, A1, A2, A3, R0 any](f func(a0 A0, a1 A1, a2 A2, a3 A3) R0) func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) R0 {
	return func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) R0 {
		return func(a1 A1) func(a2 A2) func(a3 A3) R0 {
			return func(a2 A2) func(a3 A3) R0 {
				return func(a3 A3) R0 {
					return f(a0, a1, a2, a3)
				}
			}
		}
	}
}

func A5R1[A0, A1, A2, A3, A4, R0 any](f func(a0 A0, a1 A1, a2 A2, a3 A3, a4 A4) R0) func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) R0 {
	return func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) R0 {
		return func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) R0 {
			return func(a2 A2) func(a3 A3) func(a4 A4) R0 {
				return func(a3 A3) func(a4 A4) R0 {
					return func(a4 A4) R0 {
						return f(a0, a1, a2, a3, a4)
					}
				}
			}
		}
	}
}

func A6R1[A0, A1, A2, A3, A4, A5, R0 any](f func(a0 A0, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R0) func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) R0 {
	return func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) R0 {
		return func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) R0 {
			return func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) R0 {
				return func(a3 A3) func(a4 A4) func(a5 A5) R0 {
					return func(a4 A4) func(a5 A5) R0 {
						return func(a5 A5) R0 {
							return f(a0, a1, a2, a3, a4, a5)
						}
					}
				}
			}
		}
	}
}

func A7R1[A0, A1, A2, A3, A4, A5, A6, R0 any](f func(a0 A0, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R0) func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) R0 {
	return func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) R0 {
		return func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) R0 {
			return func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) R0 {
				return func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) R0 {
					return func(a4 A4) func(a5 A5) func(a6 A6) R0 {
						return func(a5 A5) func(a6 A6) R0 {
							return func(a6 A6) R0 {
								return f(a0, a1, a2, a3, a4, a5, a6)
							}
						}
					}
				}
			}
		}
	}
}

func A8R1[A0, A1, A2, A3, A4, A5, A6, A7, R0 any](f func(a0 A0, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R0) func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) R0 {
	return func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) R0 {
		return func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) R0 {
			return func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) R0 {
				return func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) R0 {
					return func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) R0 {
						return func(a5 A5) func(a6 A6) func(a7 A7) R0 {
							return func(a6 A6) func(a7 A7) R0 {
								return func(a7 A7) R0 {
									return f(a0, a1, a2, a3, a4, a5, a6, a7)
								}
							}
						}
					}
				}
			}
		}
	}
}

func A9R1[A0, A1, A2, A3, A4, A5, A6, A7, A8, R0 any](f func(a0 A0, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R0) func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) R0 {
	return func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) R0 {
		return func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) R0 {
			return func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) R0 {
				return func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) R0 {
					return func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) R0 {
						return func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) R0 {
							return func(a6 A6) func(a7 A7) func(a8 A8) R0 {
								return func(a7 A7) func(a8 A8) R0 {
									return func(a8 A8) R0 {
										return f(a0, a1, a2, a3, a4, a5, a6, a7, a8)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func A10R1[A0, A1, A2, A3, A4, A5, A6, A7, A8, A9, R0 any](f func(a0 A0, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R0) func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) func(a9 A9) R0 {
	return func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) func(a9 A9) R0 {
		return func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) func(a9 A9) R0 {
			return func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) func(a9 A9) R0 {
				return func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) func(a9 A9) R0 {
					return func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) func(a9 A9) R0 {
						return func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) func(a9 A9) R0 {
							return func(a6 A6) func(a7 A7) func(a8 A8) func(a9 A9) R0 {
								return func(a7 A7) func(a8 A8) func(a9 A9) R0 {
									return func(a8 A8) func(a9 A9) R0 {
										return func(a9 A9) R0 {
											return f(a0, a1, a2, a3, a4, a5, a6, a7, a8, a9)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func A2R2[A0, A1, R0, R1 any](f func(a0 A0, a1 A1) (R0, R1)) func(a0 A0) func(a1 A1) (R0, R1) {
	return func(a0 A0) func(a1 A1) (R0, R1) {
		return func(a1 A1) (R0, R1) {
			return f(a0, a1)
		}
	}
}

func A3R2[A0, A1, A2, R0, R1 any](f func(a0 A0, a1 A1, a2 A2) (R0, R1)) func(a0 A0) func(a1 A1) func(a2 A2) (R0, R1) {
	return func(a0 A0) func(a1 A1) func(a2 A2) (R0, R1) {
		return func(a1 A1) func(a2 A2) (R0, R1) {
			return func(a2 A2) (R0, R1) {
				return f(a0, a1, a2)
			}
		}
	}
}

func A4R2[A0, A1, A2, A3, R0, R1 any](f func(a0 A0, a1 A1, a2 A2, a3 A3) (R0, R1)) func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) (R0, R1) {
	return func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) (R0, R1) {
		return func(a1 A1) func(a2 A2) func(a3 A3) (R0, R1) {
			return func(a2 A2) func(a3 A3) (R0, R1) {
				return func(a3 A3) (R0, R1) {
					return f(a0, a1, a2, a3)
				}
			}
		}
	}
}

func A5R2[A0, A1, A2, A3, A4, R0, R1 any](f func(a0 A0, a1 A1, a2 A2, a3 A3, a4 A4) (R0, R1)) func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) (R0, R1) {
	return func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) (R0, R1) {
		return func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) (R0, R1) {
			return func(a2 A2) func(a3 A3) func(a4 A4) (R0, R1) {
				return func(a3 A3) func(a4 A4) (R0, R1) {
					return func(a4 A4) (R0, R1) {
						return f(a0, a1, a2, a3, a4)
					}
				}
			}
		}
	}
}

func A6R2[A0, A1, A2, A3, A4, A5, R0, R1 any](f func(a0 A0, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R0, R1)) func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) (R0, R1) {
	return func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) (R0, R1) {
		return func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) (R0, R1) {
			return func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) (R0, R1) {
				return func(a3 A3) func(a4 A4) func(a5 A5) (R0, R1) {
					return func(a4 A4) func(a5 A5) (R0, R1) {
						return func(a5 A5) (R0, R1) {
							return f(a0, a1, a2, a3, a4, a5)
						}
					}
				}
			}
		}
	}
}

func A7R2[A0, A1, A2, A3, A4, A5, A6, R0, R1 any](f func(a0 A0, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R0, R1)) func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) (R0, R1) {
	return func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) (R0, R1) {
		return func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) (R0, R1) {
			return func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) (R0, R1) {
				return func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) (R0, R1) {
					return func(a4 A4) func(a5 A5) func(a6 A6) (R0, R1) {
						return func(a5 A5) func(a6 A6) (R0, R1) {
							return func(a6 A6) (R0, R1) {
								return f(a0, a1, a2, a3, a4, a5, a6)
							}
						}
					}
				}
			}
		}
	}
}

func A8R2[A0, A1, A2, A3, A4, A5, A6, A7, R0, R1 any](f func(a0 A0, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R0, R1)) func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) (R0, R1) {
	return func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) (R0, R1) {
		return func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) (R0, R1) {
			return func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) (R0, R1) {
				return func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) (R0, R1) {
					return func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) (R0, R1) {
						return func(a5 A5) func(a6 A6) func(a7 A7) (R0, R1) {
							return func(a6 A6) func(a7 A7) (R0, R1) {
								return func(a7 A7) (R0, R1) {
									return f(a0, a1, a2, a3, a4, a5, a6, a7)
								}
							}
						}
					}
				}
			}
		}
	}
}

func A9R2[A0, A1, A2, A3, A4, A5, A6, A7, A8, R0, R1 any](f func(a0 A0, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R0, R1)) func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) (R0, R1) {
	return func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) (R0, R1) {
		return func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) (R0, R1) {
			return func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) (R0, R1) {
				return func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) (R0, R1) {
					return func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) (R0, R1) {
						return func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) (R0, R1) {
							return func(a6 A6) func(a7 A7) func(a8 A8) (R0, R1) {
								return func(a7 A7) func(a8 A8) (R0, R1) {
									return func(a8 A8) (R0, R1) {
										return f(a0, a1, a2, a3, a4, a5, a6, a7, a8)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func A10R2[A0, A1, A2, A3, A4, A5, A6, A7, A8, A9, R0, R1 any](f func(a0 A0, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R0, R1)) func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) func(a9 A9) (R0, R1) {
	return func(a0 A0) func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) func(a9 A9) (R0, R1) {
		return func(a1 A1) func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) func(a9 A9) (R0, R1) {
			return func(a2 A2) func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) func(a9 A9) (R0, R1) {
				return func(a3 A3) func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) func(a9 A9) (R0, R1) {
					return func(a4 A4) func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) func(a9 A9) (R0, R1) {
						return func(a5 A5) func(a6 A6) func(a7 A7) func(a8 A8) func(a9 A9) (R0, R1) {
							return func(a6 A6) func(a7 A7) func(a8 A8) func(a9 A9) (R0, R1) {
								return func(a7 A7) func(a8 A8) func(a9 A9) (R0, R1) {
									return func(a8 A8) func(a9 A9) (R0, R1) {
										return func(a9 A9) (R0, R1) {
											return f(a0, a1, a2, a3, a4, a5, a6, a7, a8, a9)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
