import requests
import json

class API():

	def __init__(self, remote):
		if remote:
			self.url = "http://shibboleth.student.rit.edu/"
		else:
			self.url = "http://localhost:3000/api/"
		self.token = ""

	def query(self,text,typ):
		r = {}
		r['text'] = text
		r['type'] = typ
		r = requests.post(self.url + "query", data=json.dumps(r), verify=False)
		data = r.json()
		print(data)

	def question(self,key):
		r = {}
		r['key'] = key
		r = requests.post(self.url + "question", data=json.dumps(r), verify=False)
		data = r.json()
		print(data)

	def queue(self):
		r = requests.get(self.url + "queue", verify=False)
		data = r.json()
		print(data)

	def answer(self, key, answer):
		r = {}
		r['key'] = key
		r['answer'] = answer
		r['token'] = self.token
		print(r)
		r = requests.post(self.url + "answer", data=json.dumps(r), verify=False)
		data = r.json()
		print(data)

	def check(self, key):
		r = {}
		r['key'] = key
		r = requests.post(self.url + "check", data=json.dumps(r), verify=False)
		data = r.json()
		print(data)

	def recent(self):
		r = requests.get(self.url + "recent", verify=False)
		data = r.json()
		print(data)

	def login(self, username, password):
		r = {}
		r['username'] = username
		r['password'] = password
		r = requests.post(self.url + "login", data=json.dumps(r), verify=False)
		if r.status_code == requests.codes.ok:
			data = r.json()
			print("Login Successful")
			self.token = data["token"]
			return True
		print("Login Failed")
		return False

	def renew(self):
		r = {}
		r['token'] = self.token
		r = requests.post(self.url + "renew", data=json.dumps(r), verify=False)
		if r.status_code == requests.codes.ok:
			data = r.json()
			print("Renew Successful")
			self.token = data["token"]
			return
		print("Renew Failed")
		return False

if __name__ == '__main__':
	print("Jimmy CLI Starting.. ")
	i = input("local or remote? ")
	if i == "remote":
		remote = True
	else:
		remote = False
	jimmy = API(remote)
	l = False
	while(not l):
		username = input("username> ")
		password = input("password> ")
		l = jimmy.login(username, password)
	#Start REPL
	print("Use command 'help' for more options.")
	while True:
		i = input("Enter command> ")
		if i == "query":
			text = input("query> ")
			typ = "search"
			jimmy.query(text, typ)
		elif i == "question":
			key = int(input("key> "))
			jimmy.question(key)
		elif i == "queue":
			jimmy.queue()
		elif i == "answer":
			key = int(input("key> "))
			answer = input("answer> ")
			jimmy.answer(key, answer)
		elif i == "check":
			key = int(input("key> "))
			jimmy.check(key)
		elif i == "recent":
			jimmy.recent()
		elif i == "renew":
			jimmy.renew()
		elif i == "help":
			print("Avaliable commands: help, query, question, queue, answer, check, recent")
		elif i == "quit":
			break
