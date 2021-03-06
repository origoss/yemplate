apiVersion: batch/v1
kind: Job
metadata:
  name: git-copy-1
spec:
  template:
    spec:
      containers:
      - name: git-copier
        image: busybox
        volumeMounts:
        - name: gitrepo
          mountPath: /git
        - name: target
          mountPath: /target
        command:
        - sh
        - -c
        - |
          rm -rf /target/*
          rm -rf /target/.*
          cp -a /git/* /target
          ls -a /target/*/*
      restartPolicy: Never
      volumes:
      - name: gitrepo
        gitRepo:
          repository: {{ .repository }}
      - name: target
        persistentVolumeClaim:
          claimName: git-repo-target
